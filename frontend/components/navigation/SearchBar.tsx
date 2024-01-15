"use client";
import { usePathname, useSearchParams, useRouter } from "next/navigation";
import { useDebouncedCallback } from "use-debounce";

export default function SearchBar() {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();

  const handleSearch = useDebouncedCallback((term: string) => {
    const params = new URLSearchParams(searchParams);
    if (term) {
      params.set("query", term);
    } else {
      params.delete("query");
    }
    replace(`${pathname}?${params.toString()}`);
  }, 300);

  return (
    <div className="relative flex-1 rounded-full shadow-sm">
      <div className="h-full w-8 pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3 text-black">
        <svg
          width="24"
          height="24"
          viewBox="0 0 24 24"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <g>
            <path
              d="M10.2855 19.7143C15.4928 19.7143 19.7141 15.493 19.7141 10.2857C19.7141 5.07849 15.4928 0.857178 10.2855 0.857178C5.07824 0.857178 0.856934 5.07849 0.856934 10.2857C0.856934 15.493 5.07824 19.7143 10.2855 19.7143Z"
              stroke="black"
            />
            <path d="M23.1426 23.1429L17.1426 17.1429" stroke="black" />
          </g>
        </svg>
      </div>
      <input
        className="block w-full text-base rounded-full border border-border py-3 pl-12 pr-3 text-black bg-white placeholder:text-border focus:outline-none sm:text-sm sm:leading-6"
        placeholder="Search..."
        onChange={(e) => {
          handleSearch(e.target.value);
        }}
        defaultValue={searchParams.get("query")?.toString()}
      />
    </div>
  );
}
