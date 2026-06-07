## loading.tsx

- `loading.tsx` helps create loading states that users see while waiting for content to load in a specific route segment.

- Loading states appear instantly during navigation, letting users know that the application is responsive and actively loading content.

### Example Structure

```text
app/
├── dashboard/
│   ├── loading.tsx
│   └── page.tsx
```

### Example

```tsx
export default function Loading() {
  return <h2>Loading...</h2>;
}
```

### How It Works

- When a user navigates to `/dashboard`, Next.js immediately displays `loading.tsx`.
- Once the page data is loaded, `page.tsx` replaces the loading UI.

## Benefits

### 1. Immediate User Feedback

- It gives users immediate feedback when they navigate somewhere new.
- This makes your application feel fast and responsive.
- Users know their action (click/navigation) has been registered.

### 2. Shared Layouts Remain Interactive

- Next.js keeps shared layouts interactive while new content loads.
- Users can continue using navigation menus, sidebars, and other shared UI elements even when the main content is still loading.

### Note

```text
loading.tsx → Loading UI
page.tsx    → Actual page content
```