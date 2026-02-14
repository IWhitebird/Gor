'use client'

import dynamic from 'next/dynamic'

const PlaygroundEditor = dynamic(
  () => import('@/components/PlaygroundEditor'),
  { ssr: false }
)

export default function PlaygroundPage() {
  return <PlaygroundEditor />
}
