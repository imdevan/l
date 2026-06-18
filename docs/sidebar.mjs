import config from './config.mjs';

const apiReference = {
  label: 'API Reference',
  items: [
            { label: 'app', link: '/api/app' },
            { label: 'config', link: '/api/config' },
            { label: 'domain', link: '/api/domain' },
            { label: 'errors', link: '/api/errors' },
            { label: 'package', link: '/api/package' },
            { label: 'ui', link: '/api/ui' },
            { label: 'utils', link: '/api/utils' },
            { label: 'workflow', link: '/api/workflow' },
    {
      label: 'Adapters',
      items: [
              { label: 'clipboard', link: '/api/adapters/clipboard' },
              { label: 'editor', link: '/api/adapters/editor' },
              { label: 'icon', link: '/api/adapters/icon' },
              { label: 'shell', link: '/api/adapters/shell' },
              { label: 'tty', link: '/api/adapters/tty' },
      ],
    },
  ],
};

const sidebar = [
  {
    label: 'l',
    link: '/',
  },
  {
    label: 'Install',
    link: '/install',
  },
  {
    label: 'Commands',
    items: [
      { label: 'l', link: '/commands/l' },
    ],
  },
  {
    label: 'Configuration',
    link: '/configuration',
  },
];

// Add API Reference if not in production or if this is go-cli-template
const isProduction = process.env.NODE_ENV === 'production';
const projectName = 'l';
if (!isProduction || projectName === 'go-cli-template') {
  sidebar.push(apiReference);
}

sidebar.push({ label: 'Contributing', link: '/contributing' });
export default sidebar;
