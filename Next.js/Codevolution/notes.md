- Next.js is a React framework for building full-stack web applications.

# React
- It's not feasible to create a fully-featured application ready for production responsible only for the view-layer of an application.
- React is a library for building user interfaces.
- You need to make decisions about other features such as routing, data fetching, and more.

# Next.js
* It uses React for building user interfaces.
* Provides additional features that enable you to build production-ready applications.
* These features include routing, optimized rendering, data fetching, bundling, compiling, and more.
* You don't need to install additional packages as Next.js provides everything you need.
* Opinions and conventions should be followed to implement these features.

# why learn Next.js
Next.js simplifies the process of buil
- Routing (file-based routing)
- API routes (frontend + backend in one application allowing seamless integration between your frontend and backend code)
- Rendering (client-side and server side, if implemented properly, provides better performance and better search engine optimization)
- Streamlined Data fetching (nextjs provides built-in async support in react components making fetching easy)
- Styling (supports css, tailwind)
- Optimization (provides optimization for images, fonts and scripts)
- Dev and prod build system (it hepls user to not worry about configuration)

# Project structure
- run "npm run dev" -> execution starts from package.json, moves towards layout.tsx rendering the root layout component.

# React Server Components (RSC)

* React Server Components is a new architecture that was introduced by the React team and quickly adopted by Next.js.
* This architecture introduces a new approach to creating React components by dividing them into two distinct types:

  * Server components
  * Client components.

### Server Components

* By default, Next.js treats all components as Server components.
* These components can perform server-side tasks like reading files or fetching data directly from a database.
* The trade-off is that they can't use React hooks or handle user interactions.
* In this, we can use `async await` to resolve the promise and access the dynamic segments.

### Client Components

* To create a Client component, you'll need to add the `"use client"` directive at the top of your component file.
* While Client components can't perform server-side tasks like reading files, they can use hooks and handle user interactions.
* Client components are the traditional React components you're already familiar with from previous versions of React.

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