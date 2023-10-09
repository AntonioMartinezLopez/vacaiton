'use client'

import { motion, MotionValue, useInView, useScroll, useSpring, useTransform } from 'framer-motion';
import dynamic from 'next/dynamic';
import React, { useEffect, useRef, useState } from 'react'

const TestGlobe = dynamic(() => import('./testGlobe'), {
  ssr: false
})

const trips: Record<string, object[]> = {
  "europe": [{
    startLat: "52.520008",
    startLng: "13.404954",
    endLat: "48.864716",
    endLng: "2.349014",
  }, {
    startLat: "48.864716",
    startLng: "2.349014",
    endLat: "40.416775",
    endLng: "-3.703790",
  },
  {
    startLat: "41.902782",
    startLng: "-3.703790",
    endLat: "41.902782",
    endLng: "12.496366",
  },
  {
    startLat: "41.902782",
    startLng: "12.496366",
    endLat: "52.520008",
    endLng: "13.404954",
  },
  ],
  "us-east": [
    {
      startLat: "40.730610",
      startLng: "-73.935242",
      endLat: "38.889805",
      endLng: "-77.009056",
    },
    {
      startLat: "38.889805",
      startLng: "-77.009056",
      endLat: "36.174465",
      endLng: "-86.767960",
    },
    {
      startLat: "36.174465",
      startLng: "-86.767960",
      endLat: "43.092461",
      endLng: "-79.047150",
    },
    {
      startLat: "43.092461",
      startLng: "-79.047150",
      endLat: "40.730610",
      endLng: "-73.935242",
    },
  ],
  "us-west": [
    {
      startLat: "37.773972",
      startLng: "-122.431297",
      endLat: "36.188110",
      endLng: "-115.176468",
    },
    {
      startLat: "36.188110",
      startLng: "-115.176468",
      endLat: "33.448376",
      endLng: "-112.074036",
    },
    {
      startLat: "33.448376",
      startLng: "-112.074036",
      endLat: "34.052235",
      endLng: "-118.243683",
    },
    {
      startLat: "34.052235",
      startLng: "-118.243683",
      endLat: "37.773972",
      endLng: "-122.431297",
    },
  ],
  "south-america": [
    {
      startLat: "-22.908333",
      startLng: "-43.196388",
      endLat: "-30.033056",
      endLng: "-51.230000",
    },
    {
      startLat: "-30.033056",
      startLng: "-51.230000",
      endLat: "-34.603722",
      endLng: "-58.381592",
    },
    {
      startLat: "-34.603722",
      startLng: "-58.381592",
      endLat: "-25.30066",
      endLng: "-57.63591",
    },
    {
      startLat: "-25.30066",
      startLng: "-57.63591",
      endLat: "-22.908333",
      endLng: "-43.196388",
    },
  ],
  "asia-east": [
    {
      startLat: "35.652832",
      startLng: "139.839478",
      endLat: "37.532600",
      endLng: "127.024612",
    },
    {
      startLat: "37.532600",
      startLng: "127.024612",
      endLat: "31.224361",
      endLng: "121.469170",
    },
    {
      startLat: "31.224361",
      startLng: "121.469170",
      endLat: "25.105497",
      endLng: "121.597366",
    },
    {
      startLat: "25.105497",
      startLng: "121.597366",
      endLat: "35.652832",
      endLng: "139.839478",
    },
  ]
}

// function useParallax(value: MotionValue<number>, distance: number) {
//   return useTransform(value, [0, 1], [-distance, distance]);
// }

// function Section({ id }: { id: number }) {
//   const ref = useRef(null);
//   const height = useRef(0);

//   const { scrollYProgress } = useScroll({ target: ref });
//   const y = useParallax(scrollYProgress, 0);

//   return (
//     <section className='h-screen w-full flex flex-row snap-center relative'>
//       <div ref={ref} className="h-[400px] relative">
//       </div>
//       <motion.h2 style={{ y }}>{`#00${id}`}</motion.h2>

//     </section>
//   );
// }

export default function Home() {

  const [arcsData, setArcsData] = useState([] as Array<object>);
  const refContainer = React.useRef<HTMLDivElement>(null);
  const refSection1 = React.useRef<HTMLDivElement>(null);
  const refSection2 = React.useRef<HTMLDivElement>(null);
  const refSection3 = React.useRef<HTMLDivElement>(null);

  const refOneInView = useInView(refSection1);
  const refTwoInView = useInView(refSection1);
  const refThreeInView = useInView(refSection1);

  const { scrollYProgress, scrollY } = useScroll({ container: refContainer });
  const isInView = useInView(refContainer)
  const scaleX = useSpring(scrollYProgress, {
    stiffness: 100,
    damping: 30,
    restDelta: 0.001
  })

  useEffect(() => { console.log(scrollYProgress, scrollY) }, [scrollYProgress, scrollY])

  useEffect(() => { console.log(isInView) }, [isInView])

  useEffect(() => {

    if (refOneInView) {
      //focus to North America
    }
    else if (refOneInView) { }

  }, [refOneInView, refTwoInView, refThreeInView])

  useEffect(() => {
    setTimeout(() => {
      setArcsData((prev) => {
        return [...prev, ...[
          trips["europe"][0],
          trips["us-east"][0],
          trips["us-west"][0],
          trips["asia-east"][0],
          trips["south-america"][0]
        ]]
      })
    }, 1000)

    setTimeout(() => {
      setArcsData((prev) => {
        return [...prev, ...[
          trips["europe"][1],
          trips["us-east"][1],
          trips["us-west"][1],
          trips["asia-east"][1],
          trips["south-america"][1]
        ]]
      })
    }, 2000)

    setTimeout(() => {
      setArcsData((prev) => {
        return [...prev, ...[
          trips["europe"][2],
          trips["us-east"][2],
          trips["us-west"][2],
          trips["asia-east"][2],
          trips["south-america"][2]
        ]]
      })
    }, 3000)

    setTimeout(() => {
      setArcsData((prev) => {
        return [...prev, ...[
          trips["europe"][3],
          trips["us-east"][3],
          trips["us-west"][3],
          trips["asia-east"][3],
          trips["south-america"][3]
        ]]
      })
    }, 4000)




  }, [])




  return (
    <div ref={refContainer} className="relative snap-y snap-mandatory overflow-y-auto overflow-hidden h-screen scroll-smooth [&::-webkit-scrollbar]:hidden [-ms-overflow-style:'none'] [scrollbar-width:'none']">

      <section ref={refSection1} className='h-screen w-full text-gray-100 flex flex-row snap-center relative bg-gradient-to-r from-slate-800 to-slate-900 dark:bg-slate-900'>
        Welcome to Vacaiton!
        {/* <TestGlobe arcsData={arcsData} setVisible={() => null}></TestGlobe> */}
      </section>

      <section ref={refSection2} className='h-screen w-full text-gray-100 flex flex-row snap-center relative bg-gradient-to-r from-slate-800 to-slate-900 dark:bg-slate-900'>
        Welcome to Vacaiton!
      </section>

      <section ref={refSection3} className='h-screen w-full text-gray-100 flex flex-row snap-center relative bg-gradient-to-r from-slate-800 to-slate-900 dark:bg-slate-900'>
        Welcome to Vacaiton!
      </section>
      {/* <Section id={1}></Section>
      <Section id={2}></Section>
      <Section id={3}></Section> */}
      <motion.div className="fixed bottom-0 left-0 right-0 h-1 origin-[0%] bg-gradient-to-r from-cyan-600 to-cyan-700" style={{ scaleX }}></motion.div>
      <div className='fixed w-[50%] right-0 h-screen top-0 z-10 flex flex-row items-center'><TestGlobe arcsData={arcsData} setVisible={() => console.log("loaded")} pointOfView={{ lat: "39.099724", long: "-94.578331" }}></TestGlobe></div>

    </div >
  )
}
