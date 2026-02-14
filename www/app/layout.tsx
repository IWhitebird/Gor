import { RootProvider } from 'fumadocs-ui/provider'
import { Inter, JetBrains_Mono } from 'next/font/google'
import type { ReactNode } from 'react'
import './globals.css'

const inter = Inter({ subsets: ['latin'], variable: '--font-sans' })
const mono = JetBrains_Mono({ subsets: ['latin'], variable: '--font-mono' })

export const metadata = {
  title: 'Gor - An Interpreted Programming Language Written in Go',
  description: 'Gor is an interpreted programming language written in Go with JavaScript-like syntax. Zero external dependencies.',
  icons: { icon: '/gor-logo.png' },
}

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en" className={`${inter.variable} ${mono.variable}`} suppressHydrationWarning>
      <body>
        <RootProvider
          theme={{ defaultTheme: 'dark', forcedTheme: 'dark' }}
          search={{ enabled: false }}
        >
          {children}
        </RootProvider>
      </body>
    </html>
  )
}
