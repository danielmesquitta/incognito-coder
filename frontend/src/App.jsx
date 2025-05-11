import { useEffect, useState } from 'react'
import { CaptureScreenshot, GenerateSolution, Reset, SetLanguage } from '../wailsjs/go/main/App'
import { EventsOn } from '../wailsjs/runtime'
import './App.css'

function App() {
  const [solution, setSolution] = useState('')
  const [language, setLanguage] = useState('golang')
  const [isOverlayVisible, setIsOverlayVisible] = useState(false)
  const [thoughts, setThoughts] = useState('')
  const [complexity, setComplexity] = useState({ time: '', space: '' })

  useEffect(() => {
    // Listen for global keyboard shortcuts
    const unsubscribe = EventsOn("global-shortcut", async (shortcut) => {
      switch (shortcut) {
        case "screenshot":
          try {
            await CaptureScreenshot()
          } catch (error) {
            console.error('Failed to capture screenshot:', error)
          }
          break;
        
        case "generate":
          try {
            const result = await GenerateSolution()
            setSolution(result)
            setIsOverlayVisible(true)
          } catch (error) {
            console.error('Failed to generate solution:', error)
          }
          break;
        
        case "reset":
          try {
            await Reset()
            setSolution('')
            setThoughts('')
            setComplexity({ time: '', space: '' })
            setIsOverlayVisible(false)
          } catch (error) {
            console.error('Failed to reset:', error)
          }
          break;
      }
    })

    return () => {
      unsubscribe()
    }
  }, [])

  const handleLanguageChange = async (newLang) => {
    setLanguage(newLang)
    await SetLanguage(newLang)
  }

  return (
    <div className="app">
      <div className="controls">
        <select 
          value={language} 
          onChange={(e) => handleLanguageChange(e.target.value)}
          className="language-select"
        >
          <option value="golang">Go</option>
          <option value="javascript">JavaScript</option>
          <option value="java">Java</option>
          <option value="python">Python</option>
        </select>
      </div>

      {isOverlayVisible && (
        <div className="solution-overlay">
          <div className="solution-header">
            <h3>Solution</h3>
            <button onClick={() => setIsOverlayVisible(false)}>×</button>
          </div>
          <div className="solution-content">
            <pre className="code-block">{solution}</pre>
            <div className="thoughts">
              <h4>Thoughts</h4>
              <p>{thoughts}</p>
            </div>
            <div className="complexity">
              <h4>Complexity Analysis</h4>
              <p>Time Complexity: {complexity.time}</p>
              <p>Space Complexity: {complexity.space}</p>
            </div>
          </div>
        </div>
      )}
    </div>
  )
}

export default App
