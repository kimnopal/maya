import AuthTemplate from "@/components/template/AuthTemplate";
import Link from "next/link";
import React from "react";

const HelpPage = () => {
  return (
    <AuthTemplate>
      <div className="flex justify-between items-end pb-6">
        <div className="flex flex-col gap-1">
          <h1 className="font-bold text-xl">Welcome to Maya</h1>
          <p className="text-xs">Sign in to get started!</p>
        </div>
        <Link
          href="/"
          className="flex justify-center items-center gap-1 rounded-full shadow-md px-4 py-1"
        >
          <svg
            width="24"
            height="24"
            viewBox="0 0 54 54"
            fill="none"
            xmlns="http://www.w3.org/2000/svg"
          >
            <g>
              <path
                d="M28.212 17.2166C27.8905 16.8952 27.4546 16.7147 27 16.7147C26.5454 16.7147 26.1095 16.8952 25.788 17.2166L15.5023 27.5023C15.19 27.8256 15.0172 28.2586 15.0211 28.7081C15.025 29.1576 15.2053 29.5876 15.5232 29.9054C15.841 30.2233 16.271 30.4035 16.7204 30.4074C17.1699 30.4114 17.603 30.2386 17.9263 29.9263L18.4286 29.424V37.2857C18.4286 37.7404 18.6092 38.1764 18.9307 38.4979C19.2522 38.8194 19.6882 39 20.1428 39H33.8571C34.0823 39 34.3052 38.9557 34.5132 38.8695C34.7212 38.7834 34.9101 38.6571 35.0693 38.4979C35.2285 38.3387 35.3548 38.1497 35.4409 37.9418C35.5271 37.7338 35.5714 37.5108 35.5714 37.2857V29.424L36.0737 29.9263C36.2318 30.09 36.421 30.2206 36.6302 30.3105C36.8393 30.4003 37.0643 30.4476 37.2919 30.4496C37.5195 30.4516 37.7452 30.4082 37.9559 30.322C38.1666 30.2358 38.358 30.1085 38.519 29.9475C38.6799 29.7866 38.8072 29.5952 38.8934 29.3845C38.9796 29.1738 39.023 28.9481 39.021 28.7205C39.019 28.4928 38.9717 28.2679 38.8819 28.0587C38.792 27.8496 38.6614 27.6604 38.4977 27.5023L28.212 17.2166Z"
                fill="black"
              />
            </g>
          </svg>
          Back Home
        </Link>
      </div>
    </AuthTemplate>
  );
};

export default HelpPage;
