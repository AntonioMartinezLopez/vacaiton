import React from "react";
import { useEffect, useState, useRef, useCallback } from "react";
import HEX_DATA from "./data/countries_hex_data.json";
import Globe from "react-globe.gl";
import { SpotLight } from "three";

export default function CustomGlobe({ volcanoes }: any) {

    const globeEl = useRef<any>();
    const [hex, setHex] = useState<any>({ features: [] });

    useEffect(() => {
        setTimeout(() => {
            const directionalLight = globeEl.current
                .scene()
                .children.find((obj3d: { type: string }) => obj3d.type === 'DirectionalLight');
            directionalLight && directionalLight.position.set(0, 0, 0);
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
        const aLight = new SpotLight(0xffffff, 0);
        aLight.position.set(75, 500, 0);
        camera.add(aLight);

        globe.scene().add(camera);
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
            atmosphereColor={"#553C9A"}
            ref={globeEl}
            width={600}
            height={600}
            waitForGlobeReady={true}
            atmosphereAltitude={0.15}
            showGlobe={false}
            // globeImageUrl="//unpkg.com/three-globe/example/img/earth-night.jpg"
            pointsData={volcanoes}
            pointLat="lat"
            pointLng="lon"
            // COUNTRIES
            hexPolygonsData={hex.features}
            hexPolygonResolution={useCallback(() => 3, [])} //values higher than 3 makes it buggy
            hexPolygonMargin={useCallback(() => 0.6, [])} // you can mess with this to see smaller or bigger dots
            hexPolygonColor={useCallback(() => "#EE2A8F", [])}
            hexPolygonCurvatureResolution={useCallback(() => 7, [])}
        />
    );
}
