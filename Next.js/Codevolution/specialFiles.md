## Special Files in Next.js App Router

### `page.tsx`
- Defines the UI for a route.
- Required to make a route publicly accessible.

### `layout.tsx`
- Defines shared UI for a route segment and its children.
- Preserves state and remains mounted during navigation.

### `template.tsx`
- Similar to a layout but creates a fresh instance on every navigation.
- Resets state and re-runs effects.

### `loading.tsx`
- Displays a loading UI while route content is being fetched or rendered.
- Appears instantly during navigation.

### `not-found.tsx`
- Renders a custom 404 page when a route is not found.
- Can also be triggered programmatically using `notFound()`.

### `error.tsx`
- Handles runtime errors in a route segment.
- Displays a fallback UI instead of crashing the entire application.
- Must be a Client Component (`"use client"`).


## Next.js Component Hierarchy

```text
<Layout>
  <Template>
    <ErrorBoundary fallback={<Error />}>
      <Suspense fallback={<Loading />}>
        <ErrorBoundary fallback={<NotFound />}>
          <Page />
        </ErrorBoundary>
      </Suspense>
    </ErrorBoundary>
  </Template>
</Layout>
```