"use client";
import { useRouter } from "next/navigation";
import React from "react";

const BackButton = () => {
  const router = useRouter();

  return (
    <button
      onClick={router.back}
      className="bg-white border border-border rounded-full p-3"
    >
      <svg
        width="24"
        height="24"
        viewBox="0 0 24 24"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <g>
          <path d="M23.1429 12L0.857147 12" stroke="black" />
          <path d="M6.85715 6L0.857147 12L6.85715 18" stroke="black" />
        </g>
        <defs>
          <clipPath id="clip0_178_410">
            <rect
              width="24"
              height="24"
              fill="white"
              transform="matrix(0 -1 1 0 0 24)"
            />
          </clipPath>
        </defs>
      </svg>
    </button>
  );
};

export default BackButton;
