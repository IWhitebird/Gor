'use client'

import { useEffect, useRef } from 'react'

function highlightGor(code: string): string {
  return code
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/(#.*)/g, '<span class="gor-comment">$1</span>')
    .replace(/("(?:[^"\\]|\\.)*")/g, '<span class="gor-string">$1</span>')
    .replace(/\b(fn|let|const|if|else|for|return)\b/g, '<span class="gor-keyword">$1</span>')
    .replace(/\b(true|false|null)\b/g, '<span class="gor-boolean">$1</span>')
    .replace(/\b(print|swap)\b/g, '<span class="gor-builtin">$1</span>')
    .replace(/\b(\d+)\b/g, '<span class="gor-number">$1</span>')
}

const SHOWCASE_CODE = `# Counter factory using closures
fn makeCounter(start) {
    let count = start
    fn increment() {
        count = count + 1
        return count
    }
    return increment
}

let counter = makeCounter(0)
print(counter())   # 1
print(counter())   # 2
print(counter())   # 3

# Each counter has its own state
let other = makeCounter(10)
print(other())     # 11
print(counter())   # 4  (independent)`

export function CodeShowcase() {
  const codeRef = useRef<HTMLElement>(null)

  useEffect(() => {
    if (codeRef.current) {
      codeRef.current.innerHTML = highlightGor(SHOWCASE_CODE)
    }
  }, [])

  return (
    <section className="code-showcase">
      <div className="section-header">
        <h2>See Gor in Action</h2>
        <p>Clean, expressive syntax for algorithms, data structures, and functional patterns.</p>
      </div>
      <div className="code-block">
        <div className="code-block-header">
          <span className="code-block-dot"></span>
          <span className="code-block-dot"></span>
          <span className="code-block-dot"></span>
          <span className="code-block-title">closure_counter.gor</span>
        </div>
        <pre><code ref={codeRef}>{SHOWCASE_CODE}</code></pre>
      </div>
    </section>
  )
}
