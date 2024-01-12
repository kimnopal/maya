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
        <g clip-path="url(#clip0_178_410)">
          <path
            d="M23.1429 12L0.857147 12"
            stroke="black"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
          <path
            d="M6.85715 6L0.857147 12L6.85715 18"
            stroke="black"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
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
