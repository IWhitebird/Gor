'use client'

import { useState, useRef, useEffect, useCallback } from 'react'
import CodeMirror from '@uiw/react-codemirror'
import { gorTheme } from '@/lib/editorTheme'
import { gorLanguage } from '@/lib/gorLang'
import { EXAMPLES } from '@/lib/examples'

export default function PlaygroundEditor() {
  const [code, setCode] = useState(EXAMPLES['Hello World'])
  const [output, setOutput] = useState('')
  const [ast, setAst] = useState('')
  const [activeTab, setActiveTab] = useState('output')
  const [isRunning, setIsRunning] = useState(false)
  const [hasError, setHasError] = useState(false)

  const panelsRef = useRef<HTMLDivElement>(null)
  const editorPanelRef = useRef<HTMLDivElement>(null)
  const outputPanelRef = useRef<HTMLDivElement>(null)

  const handleRun = useCallback(async () => {
    if (isRunning || !code.trim()) return

    setIsRunning(true)
    setOutput('')
    setAst('')
    setHasError(false)

    try {
      const res = await fetch('/api/run', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ code }),
      })
      const data = await res.json()

      if (data.error) {
        setOutput(data.error)
        setHasError(true)
        setActiveTab('output')
      } else {
        setOutput(data.output || '(no output)')
        setAst(data.ast || '')
      }
    } catch (err: any) {
      setOutput('Network error: ' + err.message)
      setHasError(true)
      setActiveTab('output')
    } finally {
      setIsRunning(false)
    }
  }, [code, isRunning])

  const handleRunRef = useRef(handleRun)
  handleRunRef.current = handleRun

  useEffect(() => {
    const handler = (e: KeyboardEvent) => {
      if ((e.ctrlKey || e.metaKey) && e.key === 'Enter') {
        e.preventDefault()
        handleRunRef.current()
      }
    }
    window.addEventListener('keydown', handler)
    return () => window.removeEventListener('keydown', handler)
  }, [])

  const handleDividerMouseDown = (e: React.MouseEvent) => {
    e.preventDefault()
    document.body.style.cursor = 'col-resize'
    document.body.style.userSelect = 'none'

    const onMouseMove = (e: MouseEvent) => {
      if (!panelsRef.current) return
      const rect = panelsRef.current.getBoundingClientRect()
      const pct = ((e.clientX - rect.left) / rect.width) * 100
      const clamped = Math.max(20, Math.min(80, pct))
      if (editorPanelRef.current) editorPanelRef.current.style.flex = `0 0 ${clamped}%`
      if (outputPanelRef.current) outputPanelRef.current.style.flex = `0 0 ${100 - clamped}%`
    }

    const onMouseUp = () => {
      document.removeEventListener('mousemove', onMouseMove)
      document.removeEventListener('mouseup', onMouseUp)
      document.body.style.cursor = ''
      document.body.style.userSelect = ''
    }

    document.addEventListener('mousemove', onMouseMove)
    document.addEventListener('mouseup', onMouseUp)
  }

  const handleExampleChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
    const name = e.target.value
    if (name && EXAMPLES[name]) {
      setCode(EXAMPLES[name])
      e.target.value = ''
    }
  }

  const renderContent = () => {
    if (activeTab === 'output') {
      if (!output) return <span className="output-placeholder">Run your code to see output here...</span>
      return output
    } else {
      if (!ast) return <span className="output-placeholder">Run your code to see the AST here...</span>
      try {
        return JSON.stringify(JSON.parse(ast), null, 2)
      } catch {
        return ast
      }
    }
  }

  return (
    <div className="pg">
      <div className="pg-toolbar">
        <div className="pg-toolbar-left">
          <div className="pg-logo">
            <img src="/gor-logo.png" alt="Gor" width={20} height={20} />
            <span>Playground</span>
          </div>
          <div className="pg-separator" />
          <div className="pg-examples">
            <select onChange={handleExampleChange} defaultValue="" aria-label="Load example">
              <option value="">Examples</option>
              {Object.keys(EXAMPLES).map(name => (
                <option key={name} value={name}>{name}</option>
              ))}
            </select>
            <svg width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2.5"><polyline points="6 9 12 15 18 9"/></svg>
          </div>
        </div>
        <div className="pg-toolbar-right">
          <span className="pg-shortcut">
            <kbd>Ctrl</kbd>
            <span>+</span>
            <kbd>Enter</kbd>
          </span>
          <button className="pg-run" onClick={handleRun} disabled={isRunning}>
            {isRunning ? (
              <><span className="pg-spinner" /> Running</>
            ) : (
              <>
                <svg width="14" height="14" fill="currentColor" viewBox="0 0 24 24"><polygon points="5 3 19 12 5 21 5 3"/></svg>
                Run
              </>
            )}
          </button>
        </div>
      </div>

      <div className="pg-panels" ref={panelsRef}>
        <div className="pg-panel pg-editor" ref={editorPanelRef}>
          <div className="pg-panel-bar">
            <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="16 18 22 12 16 6"/><polyline points="8 6 2 12 8 18"/></svg>
            <span>main.gor</span>
          </div>
          <div className="pg-panel-body">
            <CodeMirror
              value={code}
              height="100%"
              theme={gorTheme}
              extensions={[gorLanguage]}
              onChange={(value) => setCode(value)}
              style={{ height: '100%' }}
            />
          </div>
        </div>

        <div className="pg-divider" onMouseDown={handleDividerMouseDown} />

        <div className="pg-panel pg-output" ref={outputPanelRef}>
          <div className="pg-panel-bar">
            <div className="pg-tabs">
              <button className={`pg-tab${activeTab === 'output' ? ' active' : ''}`} onClick={() => setActiveTab('output')}>
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><polyline points="4 17 10 11 4 5"/><line x1="12" y1="19" x2="20" y2="19"/></svg>
                Output
              </button>
              <button className={`pg-tab${activeTab === 'ast' ? ' active' : ''}`} onClick={() => setActiveTab('ast')}>
                <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><circle cx="12" cy="5" r="3"/><line x1="12" y1="8" x2="12" y2="14"/><circle cx="6" cy="19" r="3"/><circle cx="18" cy="19" r="3"/><line x1="12" y1="14" x2="6" y2="16"/><line x1="12" y1="14" x2="18" y2="16"/></svg>
                AST
              </button>
            </div>
            <button className="pg-clear" onClick={() => { setOutput(''); setAst(''); setHasError(false) }}>
              <svg width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" strokeWidth="2"><path d="M3 6h18"/><path d="M8 6V4h8v2"/><path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6"/></svg>
              Clear
            </button>
          </div>
          <div className="pg-panel-body">
            <div className={`pg-content${hasError ? ' error' : ''}`}>
              {renderContent()}
            </div>
          </div>
        </div>
      </div>
    </div>
  )
}
