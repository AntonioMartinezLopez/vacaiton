import React, { use } from "react";
import { useEffect, useState, useRef, useCallback } from "react";
import HEX_DATA from "./data/countries_hex_data.json";
import Globe from "react-globe.gl";
import { AmbientLight, Color, DirectionalLight, MeshPhongMaterial, SpotLight } from "three";
import { OrbitControls } from "three/examples/jsm/controls/OrbitControls";

export default function CustomGlobe() {

    const globeEl = useRef<any>();
    const ORDER_UPDATE_INTERVAL = 1000;
    const [hex, setHex] = useState<any>({ features: [] });
    const [rendered, isRendered] = useState(false);
    const [showGlobe, setShowGlobe] = useState(false);
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

        setInterval(() => {
            console.log("calling")
            if (volcanoes.length === 1) {
                setVolcanoes([...volcanoes, {
                    "name": "Acamarachi",
                    "country": "Chile",
                    "type": "Stratovolcano",
                    "lat": -23.3,
                    "lon": -67.62,
                    "elevation": 6046
                }])
            } else {
                setVolcanoes([volcanoes[0]])
            }
        }, 2000)

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
            backgroundColor={"rgb(15 23 42)"}
            atmosphereColor={"#0ea5e9"}
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
            hexPolygonColor={useCallback(() => "#14b8a6", [])}
            hexPolygonCurvatureResolution={useCallback(() => 7, [])}
        />
    );
}
