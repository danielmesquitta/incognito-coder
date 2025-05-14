import { useEffect, useState } from "react";
import {
  CaptureScreenshot,
  GenerateSolution,
  Reset,
  SetLanguage,
} from "../wailsjs/go/app/App";
import { entity } from "../wailsjs/go/models";
import {
  EventsOn,
  WindowIsMinimised,
  WindowMinimise,
  WindowSetPosition,
  WindowUnminimise,
} from "../wailsjs/runtime/runtime";
import "./App.css";

function App() {
  const [solution, setSolution] = useState<entity.Solution | null>(null);
  const [language, setLanguage] = useState("golang");
  const [isLoading, setIsLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    const unsubscribe = EventsOn("global-shortcut", async (shortcut) => {
      setError(null);

      try {
        switch (shortcut) {
          case "screenshot":
            await CaptureScreenshot();
            break;

          case "generate":
            setIsLoading(true);
            const result = await GenerateSolution();
            setSolution(result);
            break;

          case "reset":
            await Reset();
            setSolution(null);
            break;

          case "toggle-visibility":
            const isMinimised = await WindowIsMinimised();
            if (isMinimised) {
              WindowUnminimise();
            } else {
              WindowMinimise();
            }
            break;

          case "move-left":
            WindowSetPosition(0, window.screen.height / 2 - 384);
            break;

          case "move-right":
            WindowSetPosition(
              window.screen.width - 1024,
              window.screen.height / 2 - 384
            );
            break;

          case "move-up":
            WindowSetPosition(window.screen.width / 2 - 512, 0);
            break;

          case "move-down":
            WindowSetPosition(
              window.screen.width / 2 - 512,
              window.screen.height - 768
            );
            break;
        }
      } catch (err) {
        console.error("Error:", err);
        setError(err instanceof Error ? err.message : "An error occurred");
      } finally {
        setIsLoading(false);
      }
    });

    return () => {
      unsubscribe();
    };
  }, []);

  async function handleLanguageChange(newLang: string) {
    try {
      setLanguage(newLang);
      await SetLanguage(newLang);
    } catch (err) {
      console.error("Failed to change language:", err);
      setError(
        err instanceof Error ? err.message : "Failed to change language"
      );
    }
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
          <option value="typescript">TypeScript</option>
          <option value="rust">Rust</option>
          <option value="cpp">C++</option>
        </select>
      </div>

      {error && (
        <div className="error-message">
          {error}
          <button onClick={() => setError(null)}>×</button>
        </div>
      )}

      <div className="solution-overlay">
        <div className="solution-header">
          <h3>Solution</h3>
        </div>

        {isLoading ? (
          <div className="loading">
            <div className="spinner"></div>
            <p>Generating solution...</p>
          </div>
        ) : (
          <div className="solution-content">
            {solution ? (
              <>
                <div className="thoughts">
                  <h4>My Thoughts</h4>
                  <p>{solution.thoughts}</p>
                </div>

                <div className="code">
                  <h4>Code Solution</h4>
                  <pre className="code-block">{solution.code}</pre>
                </div>

                <div className="complexity">
                  <h4>Complexity Analysis</h4>
                  <p>Time Complexity: {solution.time_complexity}</p>
                  <p>Space Complexity: {solution.space_complexity}</p>
                </div>
              </>
            ) : (
              <p className="no-solution">
                Press Ctrl+Alt+P to capture a screenshot, then Ctrl+Alt+Enter to
                generate a solution.
                <br />
                <br />
                Press Ctrl+Alt+R to reset the solution.
                <br />
                <br />
                Press Ctrl+Alt+V to toggle the hide/show the window.
                <br />
                <br />
                Press Ctrl+Alt+Left/Right/Up/Down to move the window around.
              </p>
            )}
          </div>
        )}
      </div>
    </div>
  );
}

export default App;
