import { useState, useEffect } from 'react'
import TaskForm from './components/TaskForm' // Import the form

interface Task {
  id: number
  title: string
  content: string
  status: string
}

function App() {
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState(true)

  // Move fetch logic to a function so we can reuse it
  const fetchTasks = () => {
    fetch('http://localhost:4000/v1/tasks')
      .then(response => response.json())
      .then(data => {
        setTasks(data.tasks || [])
        setLoading(false)
      })
      .catch(err => {
        console.error("Failed to fetch tasks:", err)
        setLoading(false)
      })
  }

  // Fetch on initial load
  useEffect(() => {
    fetchTasks()
  }, [])

  return (
    <div className="min-h-screen bg-gray-100 p-8 font-sans">
      <div className="max-w-2xl mx-auto">
        <header className="mb-8">
          <h1 className="text-3xl font-bold text-gray-800">Task Tracker</h1>
        </header>

        {/* 👇 Add the Form Here */}
        <TaskForm onTaskCreated={fetchTasks} />

        <div className="space-y-4">
          {loading ? (
             <p className="text-gray-500 text-center">Loading...</p>
          ) : tasks.length === 0 ? (
             <div className="text-center py-10">
               <p className="text-gray-500">No tasks yet. Add one above!</p>
             </div>
          ) : (
            tasks.map(task => (
              <div key={task.id} className="bg-white p-6 rounded-xl shadow-sm border-l-4 border-blue-500">
                <div className="flex justify-between items-start">
                  <div>
                    <h3 className="font-semibold text-lg text-gray-900">{task.title}</h3>
                    <p className="text-gray-600 mt-1">{task.content}</p>
                  </div>
                  <span className={`px-2 py-1 rounded text-xs font-bold uppercase ${
                    task.status === 'completed' ? 'bg-green-100 text-green-700' : 'bg-yellow-100 text-yellow-700'
                  }`}>
                    {task.status}
                  </span>
                </div>
              </div>
            ))
          )}
        </div>
      </div>
    </div>
  )
}

export default App