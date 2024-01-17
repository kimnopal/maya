"use client";
import { usePathname, useRouter, useSearchParams } from "next/navigation";
import React, { useState } from "react";

const category = [
  {
    name: "all",
    value: "",
  },
  {
    name: "project",
    value: "project",
  },
  {
    name: "competition",
    value: "competition",
  },
];

const CategoryList = ({ category }: { category: any }) => {
  const searchParams = useSearchParams();
  const pathname = usePathname();
  const { replace } = useRouter();

  const handleCategory = (term: string) => {
    const params = new URLSearchParams(searchParams);
    if (term) {
      params.set("category", term);
    } else {
      params.delete("category");
    }
    replace(`${pathname}?${params.toString()}`);
  };

  return (
    <button
      onClick={() => handleCategory(category.value)}
      className={`text-white text-xs px-9 py-1 ${
        searchParams.get("category")
          ? searchParams.get("category") == category.value
            ? "bg-primary rounded-t-2xl"
            : "bg-secondary rounded-full"
          : "" == category.value
          ? "bg-primary rounded-t-2xl"
          : "bg-secondary rounded-full"
      }`}
    >
      {category.name}
    </button>
  );
};

const CateogryBar = () => {
  return (
    <div className="flex flex-row gap-2 py-2">
      {category.map((category, index) => (
        <CategoryList category={category} key={index} />
      ))}
    </div>
  );
};

export default CateogryBar;
