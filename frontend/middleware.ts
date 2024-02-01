import { NextRequest, NextResponse } from "next/server";


const jwt = require("jsonwebtoken")

export function middleware(request: NextRequest) {
    
    const token = request.cookies.get('token')?.value;
    if (!token) {
        return NextResponse.redirect(new URL("/signin", request.url))
    }


    // return NextResponse.redirect(new URL("/signin", request.url))
    
        // NextResponse.redirect(new URL("/", request.url))
    // jwt.verify(auth)

    return NextResponse.next()
}
export const config = {
    matcher: ['/explore'],
  }