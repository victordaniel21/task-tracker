import { useState, useEffect } from 'react'

// 1. Define the Shape of a Task (Matches Go Struct)
interface Task {
  id: number
  title: string
  content: string
  status: string
}

function App() {
  // 2. State to hold the tasks
  const [tasks, setTasks] = useState<Task[]>([])
  const [loading, setLoading] = useState(true)

  // 3. Fetch data on load
  useEffect(() => {
    fetch('http://localhost:4000/v1/tasks')
      .then(response => response.json())
      .then(data => {
        // The API returns { "tasks": [...] }
        setTasks(data.tasks || []) 
        setLoading(false)
      })
      .catch(err => {
        console.error("Failed to fetch tasks:", err)
        setLoading(false)
      })
  }, [])

  return (
    <div className="min-h-screen bg-gray-100 p-8 font-sans">
      <div className="max-w-2xl mx-auto">
        <header className="mb-8">
          <h1 className="text-3xl font-bold text-gray-800">Task Tracker</h1>
          <div className="flex items-center gap-2 mt-2">
            <span className="px-3 py-1 bg-blue-100 text-blue-800 rounded-full text-sm font-medium">
              Backend: Connected
            </span>
          </div>
        </header>

        <div className="space-y-4">
          {loading ? (
             <p className="text-gray-500 text-center">Loading...</p>
          ) : tasks.length === 0 ? (
             <div className="bg-white p-6 rounded-xl shadow-sm text-center">
               <p className="text-gray-500">No tasks found. Create one via CURL!</p>
             </div>
          ) : (
            // 4. Render the List
            tasks.map(task => (
              <div key={task.id} className="bg-white p-6 rounded-xl shadow-sm border-l-4 border-blue-500 hover:shadow-md transition-shadow">
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