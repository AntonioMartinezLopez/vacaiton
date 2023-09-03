'use client'

import dynamic from 'next/dynamic';
import React, { useRef, useLayoutEffect, useEffect, useState } from 'react'

const TestGlobe = dynamic(() => import('./testGlobe'), {
  ssr: false
})

export default function Home() {

  const [volcanoes, setVolcanoes] = useState<Array<any>>([
    {
      "name": "Abu",
      "country": "Japan",
      "type": "Shield",
      "lat": 34.5,
      "lon": 131.6,
      "elevation": 641
    },
    {
      "name": "Acamarachi",
      "country": "Chile",
      "type": "Stratovolcano",
      "lat": -23.3,
      "lon": -67.62,
      "elevation": 6046
    }])

  useEffect(() => {

    setInterval(() => {

      setVolcanoes((prev) => {
        console.log(prev);
        if (prev.length === 1) {
          return [...prev, {
            "name": "Acamarachi",
            "country": "Chile",
            "type": "Stratovolcano",
            "lat": -23.3,
            "lon": -67.62,
            "elevation": 6046
          }]
        } else {
          return [prev[0]]
        }
      })
      console.log(volcanoes)
    }, 2000)
  }, []);


  return (
    <>
      <div className="h-full flex flex-row">
        Welcome to Vacaiton!
        <TestGlobe volcanoes={volcanoes}></TestGlobe>
      </div>
    </>
  )
}
