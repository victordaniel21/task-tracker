import { useTasks } from './hooks/useTasks';
import TaskForm from './components/TaskForm';

function App() {
  const { tasks, loading, error, deleteTask, toggleStatus, refresh } = useTasks();

  // We wrap refresh to match TaskForm's expected prop signature if needed, 
  // or you can update TaskForm to use the addTask method from the hook too!
  // For now, let's keep passing refresh to TaskForm.

  if (loading) return <div className="text-center p-10">Loading...</div>;

  return (
    <div className="min-h-screen bg-gray-100 p-8 font-sans">
      <div className="max-w-2xl mx-auto">
        <header className="mb-8">
            <h1 className="text-3xl font-bold text-gray-800">Task Tracker</h1>
        </header>

        {/* You can eventually move the Create logic into the hook too */}
        <TaskForm onTaskCreated={refresh} />

        <div className="space-y-4">
            {error && <div className="text-red-500 text-center">{error}</div>}
            
            {tasks.length === 0 ? (
                <div className="text-center py-10 text-gray-500">No tasks yet.</div>
            ) : (
                tasks.map(task => (
                    <div key={task.id} className={`bg-white p-4 rounded-xl shadow-sm border-l-4 ${
                        task.status === 'completed' ? 'border-green-500 opacity-75' : 'border-blue-500'
                    }`}>
                        <div className="flex justify-between items-start">
                            <div className="flex items-start gap-3">
                                <input
                                    type="checkbox"
                                    checked={task.status === 'completed'}
                                    onChange={() => toggleStatus(task)}
                                    className="mt-1.5 h-5 w-5 text-blue-600 rounded cursor-pointer"
                                />
                                <div>
                                    <h3 className={`font-semibold text-lg ${
                                        task.status === 'completed' ? 'line-through text-gray-400' : ''
                                    }`}>
                                        {task.title}
                                    </h3>
                                    <p className="text-gray-600 mt-1 text-sm">{task.content}</p>
                                </div>
                            </div>
                            
                            <button
                                onClick={() => {
                                    if(confirm("Delete this task?")) deleteTask(task.id)
                                }}
                                className="text-gray-400 hover:text-red-500 transition-colors"
                            >
                                <svg xmlns="http://www.w3.org/2000/svg" className="h-5 w-5" viewBox="0 0 20 20" fill="currentColor">
                                    <path fillRule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 000-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clipRule="evenodd" />
                                </svg>
                            </button>
                        </div>
                    </div>
                ))
            )}
        </div>
      </div>
    </div>
  );
}

export default App;