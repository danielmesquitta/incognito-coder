import { useEffect, useState } from "react";
import {
  CaptureScreenshot,
  GenerateSolution,
  Reset,
  SetLanguage,
} from "../wailsjs/go/app/App";
import { entity } from "../wailsjs/go/models";
import { EventsOn } from "../wailsjs/runtime/runtime";
import "./App.css";

function App() {
  const [solution, setSolution] = useState<entity.Solution | null>(null);
  const [language, setLanguage] = useState("golang");
  const [isOverlayVisible, setIsOverlayVisible] = useState(false);

  useEffect(() => {
    // Listen for global keyboard shortcuts
    const unsubscribe = EventsOn("global-shortcut", async (shortcut) => {
      switch (shortcut) {
        case "screenshot":
          try {
            await CaptureScreenshot();
          } catch (error) {
            console.error("Failed to capture screenshot:", error);
          }
          break;

        case "generate":
          try {
            const result = await GenerateSolution();
            setSolution(result);
            setIsOverlayVisible(true);
          } catch (error) {
            console.error("Failed to generate solution:", error);
          }
          break;

        case "reset":
          try {
            await Reset();
            setSolution(null);
            setIsOverlayVisible(false);
          } catch (error) {
            console.error("Failed to reset:", error);
          }
          break;
      }
    });

    return () => {
      unsubscribe();
    };
  }, []);

  async function handleLanguageChange(newLang: string) {
    setLanguage(newLang);
    await SetLanguage(newLang);
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
            <pre className="code-block">{solution?.code}</pre>
            <div className="thoughts">
              <h4>Thoughts</h4>
              <p>{solution?.thoughts}</p>
            </div>
            <div className="complexity">
              <h4>Complexity Analysis</h4>
              <p>Time Complexity: {solution?.time_complexity}</p>
              <p>Space Complexity: {solution?.space_complexity}</p>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}

export default App;
