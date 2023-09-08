import React from "react";
import { useEffect, useState, useRef, useCallback } from "react";
import HEX_DATA from "./data/countries_hex_data.json";
import Globe from "react-globe.gl";
import { AmbientLight, Color, DirectionalLight, Fog, MeshPhongMaterial, PointLight, SpotLight, TextureLoader } from "three";

// Gen random data
const N = 20;
const arcsData = [] as object[];
// for (let i = 0; i < N; i++) {
//     arcsData.push({
//         startLat: (Math.random() - 0.5) * 180,
//         startLng: (Math.random() - 0.5) * 360,
//         endLat: (Math.random() - 0.5) * 180,
//         endLng: (Math.random() - 0.5) * 360,
//         color: [['red', 'white', 'blue', 'green'][Math.round(Math.random() * 3)], ['red', 'white', 'blue', 'green'][Math.round(Math.random() * 3)]]
//     });

// }
arcsData.push({
    startLat: "52.520008",
    startLng: "13.404954",
    endLat: "40.416775",
    endLng: "-3.703790",
    color: "orange"
})

// custom globe material
const globeMaterial = new MeshPhongMaterial();
globeMaterial.bumpScale = 10;
globeMaterial.color = new Color("#160E32");
globeMaterial.emissive = new Color("#160E32");
globeMaterial.emissiveIntensity = 0.1;
globeMaterial.shininess = 0.7;

export default function CustomGlobe({ volcanoes }: any) {

    const globeEl = useRef<any>();
    const [hex, setHex] = useState<any>({ features: [] });

    useEffect(() => {
        setTimeout(() => {
            const directionalLight = globeEl.current
                .scene()
                .children.find((obj3d: { type: string }) => obj3d.type === 'DirectionalLight');
            directionalLight && directionalLight.position.set(1, 1, 1);
        });
        const globe = globeEl.current;

        setHex(HEX_DATA);

        // orbitControls
        globe.controls().autoRotate = true;
        globe.controls().autoRotateSpeed = 1;
        globe.controls().enableZoom = false;
        globe.controls().minPolarAngle = 1;
        globe.controls().maxPolarAngle = 2;

        // light & camera
        const camera = globeEl.current.camera();
        camera.aspect = window.innerWidth / window.innerHeight;
        camera.updateProjectionMatrix();
        const aLight = new AmbientLight(0xbbbbbb, 0.3)
        camera.add(aLight);
        globeEl.current.scene.background = new Color(0x040d21);

        var dLight = new DirectionalLight(0xffffff, 0.8);
        dLight.position.set(-800, 2000, 400);
        camera.add(dLight);

        var dLight1 = new DirectionalLight(0x7982f6, 0.4);
        dLight1.position.set(-200, 500, 200);
        camera.add(dLight1);

        var dLight2 = new PointLight(0x8566cc, 0.5);
        dLight2.position.set(-200, 500, 200);
        camera.add(dLight2);

        // Additional effects
        globe.scene.fog = new Fog(0x535ef3, 400, 2000);

        camera.position.z = 350;
        camera.position.x = 0;
        camera.position.y = 40;

        globe.scene().add(camera);

        console.log(globe)
    }, []);



    // useEffect(() => {
    //     // Globe Controls
    //     globeRef.current!.controls().autoRotate = true;
    //     globeRef.current!.controls().autoRotateSpeed = 1;

    //     if (globeRef.current !== undefined && window !== undefined) {
    //         const scene = globeRef.current.scene();
    //         if (scene.children.length >= 3) {
    //             // Lighting
    //             let AmbientLight: AmbientLight = scene.children[1];
    //             AmbientLight.intensity = 20;
    //             AmbientLight.castShadow = false;

    //             // let DirectionalLight: DirectionalLight = scene.children[2];
    //             // DirectionalLight.intensity = 30;
    //             // DirectionalLight.position.set(-2, 2, -2);

    //             // DirectionalLight.castShadow = false;
    //             // console.log(scene);
    //         }

    //         const controls: OrbitControls = globeRef.current.controls();
    //         if (controls) {
    //             controls.enableZoom = false;
    //             controls.maxDistance = 350;
    //         }
    //     } else {
    //         console.log("Not defined");
    //     }
    // }, [rendered]);

    return (
        <Globe
            // ENVIRONMENT
            backgroundColor={"rgba(0,0,0,0)"}
            atmosphereColor={"#3a228a"}
            ref={globeEl}
            width={600}
            height={600}
            waitForGlobeReady={true}
            atmosphereAltitude={0.35}
            showGlobe={true}
            globeMaterial={globeMaterial}
            pointsData={volcanoes}
            pointLat="lat"
            pointLng="lon"
            // COUNTRIES
            hexPolygonsData={hex.features}
            hexPolygonResolution={useCallback(() => 3, [])} //values higher than 3 makes it buggy
            hexPolygonMargin={useCallback(() => 0.6, [])} // you can mess with this to see smaller or bigger dots
            hexPolygonColor={useCallback(() => "rgba(255, 255, 255, 1)", [])}
            hexPolygonCurvatureResolution={useCallback(() => 7, [])}
            //ARCS
            arcsData={arcsData}
            arcColor={'color'}
            arcDashLength={() => 1}
            arcDashGap={() => 1}
            arcDashAnimateTime={() => 600}
            arcAltitude={() => 0.1}
            arcCircularResolution={10}
            arcStroke={() => 1}
        />
    );
}
