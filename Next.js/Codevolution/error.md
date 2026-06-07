# Errors

- `error.tsx` must always be a client component
- Error files are used so that the entire application does not break
- Wth of error files, the error will be shown only on the specified route page

# error.tsx

- `error.tsx` automatically wraps route segments and their nested children in a React Error Boundary.

- You can create custom error UIs for specific route segments using the file-system hierarchy.

- It isolates errors to affected segments while keeping the rest of the application functional.

- It enables users to recover from an error without requiring a full page reload.

### Example Structure

```text
app/
├── dashboard/
│   ├── error.tsx
│   ├── page.tsx
│   └── settings/
│       └── page.tsx
```

### Example

```tsx
"use client";

export default function Error({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
  return (
    <div>
      <h2>Something went wrong!</h2>
      <button onClick={() => reset()}>
        Try Again
      </button>
    </div>
  );
}
```

### Benefits

- Prevents the entire application from crashing.
- Displays a custom fallback UI when errors occur.
- Keeps unaffected route segments functional.
- Allows retrying failed renders using the `reset()` function.

### Note

```text
error.tsx
   ↓
React Error Boundary
   ↓
Catches errors in the current route segment
and its nested children
```

- `error.tsx` must be a Client Component, so it requires:

```tsx
"use client";
```

# Recovering from errors

1. Retry btn
- It will attempt to re-render the client side  
- the reload function will ensures that the refresh is deferred (postponed or delayed) until the next render phase allowing react to handle any pending state updates before proceeding

```tsx
'use client'

import { useRouter } from 'next/navigation'
import {startTransition} from 'react'

const ErrorBoundary = ({
    error,
    reset
} : {
    error: Error,
    reset: () => void
}) => {
    const router = useRouter()
    const reload = () => {
        startTransition(() => {
            router.refresh()
            reset()
        })
    }
  return (
    <>
    <div>{error.name}</div>
    <div>{error.message}</div>
    <button onClick={reload}>Reload</button>
    </>
  )
}

export default ErrorBoundary
```

# Handling Errors in Nested Routes

- Errors always bubble up to find the closest parent error boundary.
- An `error.tsx` file handles errors not only for its own folder, but also for all nested child route segments beneath it.
- By strategically placing `error.tsx` files at different levels in your route folders, you can control how specific or broad your error handling is.
- Where u put your error.tsx file makes a huge difference -> it determines exactly which parts of your UI get affected when things go wrong

### Example Structure

```text
app/
├── error.tsx
├── dashboard/
│   ├── error.tsx
│   ├── page.tsx
│   └── analytics/
│       └── page.tsx
```

### Error Handling Flow

```text
analytics/page.tsx throws an error
            │
            ▼
dashboard/error.tsx
```

If `dashboard/error.tsx` does not exist:

```text
analytics/page.tsx throws an error
            │
            ▼
app/error.tsx
```

### Error Bubbling

```text
analytics/page.tsx
        │
        ▼
dashboard/error.tsx
        │
        ▼
app/error.tsx
```

- Next.js looks for the nearest parent `error.tsx`.
- If none is found, the error continues bubbling upward until one is found.
- This prevents the entire application from crashing.

### Benefits

- Fine-grained error handling.
- Different error UIs for different sections of the app.
- Better user experience during failures.
- Isolates failures to specific route segments.

# Handling Errors in Layouts

- An `error.tsx` file handles errors for all of its nested child segments.
- However, there is an important limitation when working with `layout.tsx` in the same route segment.
- The error boundary **cannot catch errors thrown inside `layout.tsx` of the same segment**.

### Why?

- In the component hierarchy, `layout.tsx` sits **above** the error boundary.

```text
<Layout>
  <Template>
    <ErrorBoundary fallback={<Error />}>
      <Page />
    </ErrorBoundary>
  </Template>
</Layout>
```

- Since the error boundary is rendered *inside* the layout, it cannot catch errors originating from the layout itself.

### Example Structure

```text
app/
├── layout.tsx
├── error.tsx
└── page.tsx
```

### What Happens?

```text
layout.tsx throws an error
        │
        ▼
error.tsx ❌ NOT triggered
```

- The `error.tsx` in the same segment cannot handle errors from its own `layout.tsx`.

### How to Handle Layout Errors?

Use a parent segment's error boundary:

```text
app/
├── error.tsx
├── dashboard/
│   ├── layout.tsx
│   ├── error.tsx
│   └── page.tsx
```

If `dashboard/layout.tsx` throws an error:

```text
dashboard/layout.tsx
        │
        ▼
app/error.tsx ✅
```

### Key Point

```text
error.tsx catches:
✓ page.tsx
✓ nested layouts
✓ nested pages
✓ child route segments

error.tsx does NOT catch:
✗ layout.tsx in the same segment
```

### Rule

- Errors always bubble up to the nearest parent error boundary.
- For errors in a layout, the error boundary must exist in a parent segment.

# Handling Global Errors

- An `error.tsx` file cannot catch errors thrown in a `layout.tsx` file within the same route segment.

- This becomes a problem for the root `layout.tsx` because it has no parent segment whose error boundary can catch its errors.

- To handle these cases, Next.js provides a special file called `global-error.tsx`.

- `global-error.tsx` is placed in the root `app` directory and acts as the last line of defense for application-wide failures.

### Key Points

- Handles errors thrown in the root layout.
- Catches catastrophic errors at the highest level of the application.
- Replaces the entire application UI when triggered.
- Works only in production mode.
- Must render both `<html>` and `<body>` tags.

### Example Structure

```text
app/
├── global-error.tsx
├── layout.tsx
├── page.tsx
└── dashboard/
    └── page.tsx
```

### Example

```tsx
"use client";

export default function GlobalError({
  error,
  reset,
}: {
  error: Error;
  reset: () => void;
}) {
  return (
    <html>
      <body>
        <h2>Something went terribly wrong!</h2>
        <button onClick={() => reset()}>
          Try Again
        </button>
      </body>
    </html>
  );
}
```

### Error Handling Hierarchy

```text
global-error.tsx
        ↑
    Root Layout
        ↑
     error.tsx
        ↑
      Page
```

### Note

```text
error.tsx         → Route-level error handling
global-error.tsx  → Application-level error handling
```

- Use `error.tsx` for specific route segments.
- Use `global-error.tsx` as a fallback for unrecoverable root-level errors.