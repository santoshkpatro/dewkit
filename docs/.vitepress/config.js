import { defineConfig } from 'vitepress'

export default defineConfig({
  lang: 'en-US',
  title: 'DewKit',
  description:
    'Open-source customer support & ticketing platform built with Go, PostgreSQL, Redis, and Vue',

  lastUpdated: true,
  cleanUrls: true,

  themeConfig: {
    logo: '/logo.svg',

    nav: [
      { text: 'Docs', link: '/guide/overview' },
      { text: 'Installation', link: '/guide/installation' },
      { text: 'Deployment', link: '/guide/deployment' },
      { text: 'Roadmap', link: '/roadmap' },
      {
        text: 'Community',
        items: [
          { text: 'GitHub', link: 'https://github.com/santoshkpatro/dewkit' },
          {
            text: 'Discussions',
            link: 'https://github.com/santoshkpatro/dewkit/discussions',
          },
        ],
      },
    ],

    sidebar: {
      '/guide/': [
        {
          text: 'Introduction',
          collapsed: false,
          items: [
            { text: 'Overview', link: '/guide/overview' },
            { text: 'Why DewKit?', link: '/guide/why-dewkit' },
            { text: 'Architecture', link: '/guide/architecture' },
          ],
        },
        {
          text: 'Getting Started',
          collapsed: false,
          items: [
            { text: 'Installation', link: '/guide/installation' },
            { text: 'Configuration', link: '/guide/configuration' },
            { text: 'Running DewKit', link: '/guide/running' },
          ],
        },
        {
          text: 'Core Concepts',
          collapsed: true,
          items: [
            { text: 'Tickets', link: '/guide/tickets' },
            { text: 'Inbox & Conversations', link: '/guide/inbox' },
            { text: 'Users & Roles', link: '/guide/users-roles' },
          ],
        },
        {
          text: 'Deployment',
          collapsed: false,
          items: [
            { text: 'Deployment Overview', link: '/guide/deployment' },
            { text: 'Deploying on Heroku', link: '/guide/deploying-heroku' },
            {
              text: 'Deploying on AWS Elastic Beanstalk',
              link: '/guide/deploying-aws-eb',
            },
          ],
        },
        {
          text: 'Operations',
          collapsed: true,
          items: [
            { text: 'Operations Guide', link: '/guide/operations' },
            { text: 'Troubleshooting', link: '/guide/troubleshooting' },
          ],
        },
      ],
    },

    socialLinks: [{ icon: 'github', link: 'https://github.com/santoshkpatro/dewkit' }],

    // footer: {
    //   message: 'Released under the MIT License.',
    //   copyright: '© 2025 DewKit Contributors',
    // },

    search: {
      provider: 'local',
    },
  },
})
