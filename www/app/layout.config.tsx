import type { BaseLayoutProps } from 'fumadocs-ui/layouts/shared'
import { GitHubStars } from '@/components/GitHubStars'

export const baseOptions: BaseLayoutProps = {
  nav: {
    title: (
      <div style={{ display: 'flex', alignItems: 'center', gap: '0.5rem' }}>
        <img
          src="/gor-logo.png"
          width={28}
          height={28}
          alt="Gor"
          style={{ borderRadius: 6, objectFit: 'contain' }}
        />
        <span style={{ fontWeight: 700, fontSize: '1rem' }}>Gor</span>
        <span style={{
          fontSize: '0.65rem',
          fontWeight: 600,
          padding: '0.1rem 0.4rem',
          borderRadius: '4px',
          background: 'rgba(220, 38, 38, 0.15)',
          color: '#f87171',
          letterSpacing: '0.02em',
        }}>v1.0</span>
      </div>
    ),
  },
  links: [
    { text: 'Playground', url: '/playground' },
    { text: 'Docs', url: '/docs' },
    {
      type: 'custom',
      secondary: true,
      children: <GitHubStars />,
    },
  ],
  themeSwitch: { enabled: false },
  searchToggle: { enabled: false },
}
