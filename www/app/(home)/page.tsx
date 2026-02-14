import Link from 'next/link'
import { CodeShowcase } from '@/components/CodeShowcase'

export default function Home() {
  return (
    <>
      {/* Hero */}
      <section className="hero">
        <div className="hero-layout">
          <div className="hero-text">
            <div className="hero-badge">
              Built with <span>Go</span> &mdash; Zero Dependencies
            </div>
            <h1><span className="gradient-text">Gor</span> Programming Language</h1>
            <p className="hero-subtitle">
              An interpreted language with JavaScript-like syntax, built entirely in Go&apos;s standard library. Closures, objects, arrays &mdash; everything you need.
            </p>
            <div className="hero-actions">
              <Link href="/playground" className="btn btn-primary">
                <svg width="18" height="18" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24"><polygon points="5 3 19 12 5 21 5 3"/></svg>
                Try Playground
              </Link>
              <Link href="/docs" className="btn btn-secondary">
                <svg width="18" height="18" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24"><path d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2z"/><path d="M22 3h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"/></svg>
                Read Docs
              </Link>
            </div>
          </div>
          <div className="hero-mascot">
            <img src="/gor-logo.png" alt="Gor mascot" width={260} height={260} />
          </div>
        </div>
      </section>

      {/* Features */}
      <section className="features">
        <div className="features-grid">
          <div className="feature-card">
            <div className="feature-icon">&#9881;</div>
            <h3>Zero Dependencies</h3>
            <p>Built entirely with Go&apos;s standard library. No external packages, no bloat &mdash; just pure Go powering a complete interpreter.</p>
          </div>
          <div className="feature-card">
            <div className="feature-icon">&#123;&#125;</div>
            <h3>JS-like Syntax</h3>
            <p>Familiar let/const declarations, fn functions, for loops, and if/else blocks. Easy to pick up.</p>
          </div>
          <div className="feature-card">
            <div className="feature-icon">&lambda;</div>
            <h3>First-Class Functions</h3>
            <p>Full closure support with lexical scoping. Create counter factories, accumulators, and higher-order functions.</p>
          </div>
          <div className="feature-card">
            <div className="feature-icon">&#9638;</div>
            <h3>Objects &amp; Arrays</h3>
            <p>Rich data structures with objects (key-value pairs) and arrays. Member access, index expressions, and nested structures.</p>
          </div>
          <div className="feature-card">
            <div className="feature-icon">&#9654;</div>
            <h3>Interactive REPL</h3>
            <p>Explore the language interactively with the built-in REPL, or run source files directly from the command line.</p>
          </div>
          <div className="feature-card">
            <div className="feature-icon">&#128200;</div>
            <h3>AST Inspector</h3>
            <p>Built-in --ast flag outputs the Abstract Syntax Tree as JSON. Great for learning how interpreters work.</p>
          </div>
        </div>
      </section>

      {/* Code Showcase */}
      <CodeShowcase />

      {/* Quick Start */}
      <section className="quickstart">
        <h2>Quick Start</h2>
        <div className="quickstart-grid">
          <div className="quickstart-steps">
            <div className="quickstart-step">
              <div className="step-number">1</div>
              <div className="step-content">
                <h4>Clone &amp; Build</h4>
                <p><code>git clone https://github.com/IWhitebird/Gor.git &amp;&amp; cd Gor &amp;&amp; make build</code></p>
              </div>
            </div>
            <div className="quickstart-step">
              <div className="step-number">2</div>
              <div className="step-content">
                <h4>Run a File</h4>
                <p><code>./bin/gor examples/fibonacci.gor</code></p>
              </div>
            </div>
            <div className="quickstart-step">
              <div className="step-number">3</div>
              <div className="step-content">
                <h4>Start the REPL</h4>
                <p><code>./bin/gor --repl</code></p>
              </div>
            </div>
            <div className="quickstart-step">
              <div className="step-number">4</div>
              <div className="step-content">
                <h4>Inspect the AST</h4>
                <p><code>./bin/gor --ast examples/fibonacci.gor</code></p>
              </div>
            </div>
          </div>
          <div className="code-block">
            <div className="code-block-header">
              <span className="code-block-dot"></span>
              <span className="code-block-dot"></span>
              <span className="code-block-dot"></span>
              <span className="code-block-title">hello.gor</span>
            </div>
            <pre><code>{`# Hello World in Gor
print("Hello, World!")

# Variables
let name = "Gor"
const version = 1
print("Welcome to " + name)

# Functions
fn add(a, b) {
    return a + b
}
print(add(40, 2))`}</code></pre>
          </div>
        </div>
      </section>

      {/* Footer */}
      <footer className="site-footer">
        <p>Gor Programming Language &mdash; Written in Go</p>
        <div className="site-footer-links">
          <a href="https://github.com/IWhitebird/Gor" target="_blank" rel="noopener">GitHub</a>
          <Link href="/docs">Docs</Link>
          <Link href="/playground">Playground</Link>
        </div>
      </footer>
    </>
  )
}
