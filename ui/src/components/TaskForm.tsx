import { useState } from 'react'
import { PlusCircle } from 'lucide-react'

interface TaskFormProps {
  onTaskCreated: () => void
}

export default function TaskForm({ onTaskCreated }: TaskFormProps) {
  const [title, setTitle] = useState('')
  const [content, setContent] = useState('')
  const [isSubmitting, setIsSubmitting] = useState(false)

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!title.trim()) return // Prevent empty submissions

    setIsSubmitting(true)
    try {
      const response = await fetch('http://localhost:4000/v1/tasks', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ title, content }),
      })

      if (response.ok) {
        setTitle('')
        setContent('')
        onTaskCreated()
      }
    } catch (error) {
      console.error('Failed to create task:', error)
    } finally {
      setIsSubmitting(false)
    }
  }

  return (
    <form onSubmit={handleSubmit} className="mb-10 bg-white p-1 rounded-2xl shadow-sm border border-slate-200 focus-within:ring-2 focus-within:ring-indigo-500 focus-within:border-indigo-500 transition-all duration-200">
      <div className="p-4 border-b border-slate-100">
        <input
          type="text"
          value={title}
          onChange={(e) => setTitle(e.target.value)}
          className="w-full text-lg font-semibold text-slate-800 placeholder-slate-400 bg-transparent border-none focus:ring-0 focus:outline-none"
          placeholder="What needs to be done?"
          required
        />
      </div>
      <div className="p-4 flex items-center justify-between gap-4 bg-slate-50/50 rounded-b-2xl">
        <input
          type="text"
          value={content}
          onChange={(e) => setContent(e.target.value)}
          className="w-full text-sm text-slate-600 placeholder-slate-400 bg-transparent border-none focus:ring-0 focus:outline-none"
          placeholder="Add a quick detail or note (optional)..."
        />
        <button
          type="submit"
          disabled={isSubmitting || !title.trim()}
          className="flex items-center gap-2 bg-indigo-600 hover:bg-indigo-700 disabled:bg-slate-300 text-white px-5 py-2.5 rounded-xl font-medium transition-colors whitespace-nowrap"
        >
          <PlusCircle size={18} />
          {isSubmitting ? 'Adding...' : 'Add Task'}
        </button>
      </div>
    </form>
  )
}