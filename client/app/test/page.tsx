"use client"

import React, { useRef, useLayoutEffect, useEffect } from 'react'

import TestGlobe from './testGlobe'

export default function Home() {

  return (
    <>
      <div className="h-full bg-slate-900 flex flex-row">
        Welcome to Vacaiton!
        <TestGlobe></TestGlobe>
      </div>
    </>
  )
}
