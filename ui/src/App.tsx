import { useState, useEffect } from 'react'
import TaskForm from './components/TaskForm'
import { CheckCircle2, Circle, Trash2, LayoutList, AlertCircle } from 'lucide-react'

interface Task {
  id: number
  title: string
  content: string
  status: string
}

function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState(true)
  
  // 👇 1. New State for our custom modal
  const [taskToDelete, setTaskToDelete] = useState<number | null>(null)

  const fetchTasks = () => {
    fetch('http://localhost:4000/v1/tasks')
      .then(res => res.json())
      .then(data => {
        setTasks(data.tasks || [])
        setLoading(false)
      })
      .catch(err => console.error(err))
  }

  useEffect(() => {
    fetchTasks()
  }, [])

  const toggleStatus = async (task: Task) => {
    const newStatus = task.status === 'completed' ? 'pending' : 'completed'
    try {
      const res = await fetch(`http://localhost:4000/v1/tasks/${task.id}`, {
        method: 'PUT',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ status: newStatus }),
      })
      if (res.ok) {
        setTasks(prev => prev.map(t => t.id === task.id ? { ...t, status: newStatus } : t))
      }
    } catch (error) {
      console.error("Update failed:", error)
    }
  }

  // 👇 2. Modified Delete Function (It no longer asks for confirmation itself)
  const executeDelete = async () => {
    if (taskToDelete === null) return
    
    try {
      const res = await fetch(`http://localhost:4000/v1/tasks/${taskToDelete}`, { method: 'DELETE' })
      if (res.ok) {
        setTasks(prev => prev.filter(t => t.id !== taskToDelete))
      }
    } catch (error) {
      console.error("Delete failed:", error)
    } finally {
      // Always close the modal when done
      setTaskToDelete(null)
    }
  }

  const completedCount = tasks.filter(t => t.status === 'completed').length

  return (
    <div className="min-h-screen bg-slate-50 p-8 font-sans selection:bg-indigo-100 relative">
      <div className="max-w-2xl mx-auto">
        
        <header className="mb-10 flex justify-between items-end">
          <div>
            <h1 className="text-4xl font-extrabold text-slate-900 tracking-tight">re-Focus</h1>
            <p className="text-slate-500 mt-1 font-medium">Your personal task tracker.</p>
          </div>
          <div className="text-sm font-semibold text-slate-400 bg-white px-4 py-1.5 rounded-full shadow-sm border border-slate-200">
            {completedCount} / {tasks.length} Done
          </div>
        </header>

        <TaskForm onTaskCreated={fetchTasks} />

        <div className="space-y-3">
          {loading ? (
            <div className="flex justify-center py-12">
              <div className="animate-pulse flex items-center gap-2 text-slate-400 font-medium">
                Loading tasks...
              </div>
            </div>
          ) : tasks.length === 0 ? (
            <div className="text-center py-16 px-6 border-2 border-dashed border-slate-200 rounded-3xl bg-slate-50/50">
              <LayoutList className="mx-auto h-12 w-12 text-slate-300 mb-4" />
              <h3 className="text-lg font-semibold text-slate-700">No tasks yet</h3>
              <p className="text-slate-500 mt-1 max-w-sm mx-auto">You're all caught up! Add a new task above to get started on your day.</p>
            </div>
          ) : (
            tasks.map(task => {
              const isDone = task.status === 'completed'
              return (
                <div 
                  key={task.id} 
                  className={`group bg-white p-5 rounded-2xl border transition-all duration-300 flex justify-between items-start gap-4 ${
                    isDone 
                      ? 'border-slate-100 shadow-sm opacity-60 bg-slate-50/50' 
                      : 'border-slate-200 shadow-sm hover:shadow-md hover:border-indigo-200'
                  }`}
                >
                  <div className="flex items-start gap-4 flex-1">
                    <button 
                      onClick={() => toggleStatus(task)}
                      className="mt-0.5 text-slate-400 hover:text-indigo-600 transition-colors focus:outline-none"
                    >
                      {isDone ? <CheckCircle2 className="w-6 h-6 text-emerald-500" /> : <Circle className="w-6 h-6" />}
                    </button>
                    
                    <div className="flex-1">
                      <h3 className={`font-semibold text-lg transition-colors ${isDone ? 'line-through text-slate-500' : 'text-slate-800'}`}>
                        {task.title}
                      </h3>
                      {task.content && (
                        <p className={`mt-1 text-sm leading-relaxed ${isDone ? 'text-slate-400' : 'text-slate-600'}`}>
                          {task.content}
                        </p>
                      )}
                    </div>
                  </div>

                  {/* 👇 3. Trash icon now just opens the modal instead of deleting instantly */}
                  <button 
                    onClick={() => setTaskToDelete(task.id)}
                    className="text-slate-300 hover:text-rose-500 hover:bg-rose-50 p-2 rounded-xl transition-all opacity-0 group-hover:opacity-100 focus:opacity-100"
                    title="Delete Task"
                  >
                    <Trash2 className="w-5 h-5" />
                  </button>
                </div>
              )
            })
          )}
        </div>
      </div>

      {/* 👇 4. The Custom Delete Modal Overlay */}
      {taskToDelete !== null && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-slate-900/40 backdrop-blur-sm transition-opacity">
          <div className="bg-white rounded-2xl shadow-xl max-w-sm w-full p-6 border border-slate-100 animate-in fade-in zoom-in-95 duration-200">
            
            <div className="flex items-center gap-4 mb-4">
              <div className="bg-rose-100 p-3 rounded-full text-rose-600">
                <AlertCircle className="w-6 h-6" />
              </div>
              <h3 className="text-xl font-bold text-slate-900">Delete Task</h3>
            </div>
            
            <p className="text-slate-600 mb-8">
              Are you sure you want to delete this task? This action cannot be undone.
            </p>
            
            <div className="flex items-center justify-end gap-3">
              <button 
                onClick={() => setTaskToDelete(null)}
                className="px-5 py-2.5 rounded-xl font-medium text-slate-600 hover:bg-slate-100 transition-colors"
              >
                Cancel
              </button>
              <button 
                onClick={executeDelete}
                className="px-5 py-2.5 rounded-xl font-medium text-white bg-rose-600 hover:bg-rose-700 shadow-sm transition-colors"
              >
                Delete
              </button>
            </div>
            
          </div>
        </div>
      )}

    </div>
  )
}

export default App