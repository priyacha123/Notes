# Routing

* Next.js has a file-system based routing system.
* URLs you can access in your browser are determined by how you organize your files and folders in your code.

### Routing Conventions

1. All routes must live inside the `app` folder.
2. Route files must be named either `page.js` or `page.tsx`.
3. Each folder represents a segment of the URL path.

* When these conventions are followed, the file automatically becomes available as a route.

### Dyanamic Routing
- Every page in app router receives route parameters through the params prop.
* The type of params is a promise that results to an object containing the dynamic segments as key:value pair

```tsx
const page = async ({
    params
}: {
    params: Promise<{ id: string}>;
}) => {
    const id = (await params).id;
  return (
    <div>Product Details {id} </div>
  )
}
export default page 
```

### Nested Dyanamic Routing
```tsx
const page = async({
    params
} : {
    params: Promise<{id: string, reviewId: string}>;
}) => {
    const {id, reviewId} = await params;
  return (
    <div>Id {id} - Review Id {reviewId}</div>
  )
}

export default page
```

# Catch-all segments
- best for documentation sites where u want different URL segments for organization and SEO but keep the same basic layout

- ` ...[slug] ` is used for this
- Having multiple routes through one file
- Eg: docs/feature/12/blog
```text
app/
  |--docs/
    |-- [slug]
```

```tsx
const catchAll = async({
    params
} : {
    params : Promise<{slug: string[]}>
}) => {
    const {slug} = await params;
    if (slug.length === 2) {
        return (
            <h1>
            Viewing docs for feature {slug[0]} and second route {slug[1]}
            </h1>
        )
    }
    else if (slug.length === 1) {
        return (
            <h1>Viewing docs for feature {slug[0]} </h1>
        )
    }
    return (
    <>
<h1>Viewing all docs</h1>
    </>
  )
}

export default catchAll
```

## Not Found Page
- Customize not-found page for nested routes
```tsx
"use client";

import { usePathname } from "next/navigation";

const NotFoundRe = () => {
  const pathname = usePathname();
  const productId = pathname.split("/")[2]
  const reviewId = pathname.split("/")[4]
  return (
    <div>not-found productId {productId} - reviewId {reviewId}</div>
  )
}

export default NotFoundRe
```

## Safe collocation of file
- Only page.tsx files are public and it returns what is returned inside the `export default funName` 

## Private Folders
- A way to tell Next.js: "This folder is just for internal stuff — don't include it in the routing system."
- The folder and all its subfolders are excluded from routing.
- Add an underscore (_) at the start of the folder name.
- Example
app/
├── _components/
│   ├── Navbar.tsx
│   └── Footer.tsx
├── about/
│   └── page.tsx
└── page.tsx
- _components is a private folder and does not create a route.
- Only about/page.tsx becomes the /about route.

### Note
```text
All folders starting with _ are ignored by Next.js routing but can still be imported and used throughout the application.
```

## Private Folders
- A way to tell Next.js: **"This folder is just for internal stuff — don't include it in the routing system."**
- The folder and all its subfolders are excluded from routing.
- Add an underscore (`_`) at the start of the folder name.

### Benefits
- Keeping your UI logic separate from routing logic.
- Having a consistent way to organize internal files in your project.
- Making it easier to group related files in your code editor.
- Avoiding potential naming conflicts with future Next.js file naming conventions.

### Note
- If you actually want an underscore (`_`) in your URL, use `%5F` instead.
- `%5F` is the URL-encoded version of an underscore.

## Route Groups
- Lets us logically organize our routes and project files without impacting the URL structure.
- It also helps to apply layouts selectively to specific parts of our app
- Example
```text
app/
├── (marketing)/
│   ├── about/
│   │   └── page.tsx
│   └── contact/
│       └── page.tsx
├── (shop)/
│   ├── products/
│   │   └── page.tsx
│   └── cart/
│       └── page.tsx
└── page.tsx
```
## Nested Layout
- With the help of route groups
- Have layout.tsx file in every folder where u want customize layouts
- Make sure that page.tsx does not lie in root.
- Example
```text
app/
├── (marketing)/
│   ├── about/
│   │   └── page.tsx
│   └── contact/
│       └── page.tsx
│   └── layout.tsx
│   └── page.tsx
├── (shop)/
│   ├── products/
│   │   └── page.tsx
│   └── cart/
│       └── page.tsx
│   └── layout.tsx
│   └── page.tsx
```