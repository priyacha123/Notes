# Navigation

## Link
- By default, navigation adds a new entry to the browser history:
- use replace in `<Link>` tag to replace the current history entry instead of adding a new one.
- Example
Login
↓
Home
↓
Articles/1  (Home entry replaced)
- Now pressing Back will skip the replaced page(Home page) because it is no longer in the history stack i.e., u will go back to Login page directly, skipping Home page.

## Params and SearchParams
For a given URL:
- `params` is a promise that resolves to an object containing the dynamic route parameters (e.g., `id`).
- `searchParams` is a promise that resolves to an object containing the query parameters (e.g., filters, sorting, search terms).
- While `page.tsx` has access to both `params` and `searchParams`, `layout.tsx` only has access to `params`.

```tsx
import Link from "next/link";

const Articles = async({
  params,
  searchParams,
} : {
  params: Promise< {articleId: string}>,
  searchParams: Promise<{lang?: string}>
}) => {
  const {articleId} = await params;
  const { lang } = await searchParams;
  return (
    <div>
        <h1>Articles {articleId} </h1>
        <p>Welcome to the articles page in {lang}!</p>
        
        <div>
          <Link href={`/articles/${articleId}?lang=us`}>View Article in US </Link>
          <br />
          <Link href={`/articles/${articleId}?lang=en`}>View Article in English </Link>
          <br />
          <Link href={`/articles/${articleId}?lang=es`}>View Article in Spanish </Link>
        </div>
    </div>
  )
}

export default Articles
``` 

```text
- In a client component you do not need params and searchParams props. you can directly use useParams and useSearchParams hooks to access the dynamic route parameters and query parameters respectively.

- `useSearchParams()` does not return an object. It returns URLSearchParams, so use `.get()`.
```

