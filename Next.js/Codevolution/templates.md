# Layouts
- layouts only mount the new page content while keeping shared elements intact
- They don't remount shared components which leads to better performance


# Templates

- Templates are similar to layouts in that they are also UI shared between multiple pages in your application.

- Whenever a user navigates between routes sharing a template, you get a completely fresh start:
  - A new template component instance is mounted.
  - DOM elements are recreated.
  - Component state is cleared.
  - Effects are re-synchronized.
- Create a template by exporting a default React component from a `template.js` or `template.tsx` file.
- Like layouts, templates need to accept a `children` prop to render the nested route segments.

## Template vs Layout

| Layout | Template |
|----------|----------|
| Preserves state between navigations | Resets state on navigation |
| Component instance is reused | New component instance is mounted |
| DOM elements are preserved | DOM elements are recreated |
| Effects are not re-run unnecessarily | Effects are re-synchronized |

### Example Structure

```text
app/
├── dashboard/
│   ├── template.tsx
│   ├── analytics/
│   │   └── page.tsx
│   └── settings/
│       └── page.tsx
```

### Example Template

```tsx
export default function Template({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <div>
      <h2>Dashboard Template</h2>
      {children}
    </div>
  );
}
```

## Use Cases

- Resetting form state on navigation.
- Re-running animations between pages.
- Re-triggering effects when changing routes.
- Creating a fresh UI instance for each page visit.

### Key Points

- Templates are defined using `template.tsx` or `template.js`.
- Templates wrap all nested routes within their segment.
- Templates receive a `children` prop.
- Unlike layouts, templates remount when navigating between routes, creating a fresh component instance.
