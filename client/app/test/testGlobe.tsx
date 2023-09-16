import React from "react";
import { useEffect, useState, useRef, useCallback } from "react";
import HEX_DATA from "./data/countries_hex_data.json";
import Globe from "react-globe.gl";
import { AmbientLight, Color, DirectionalLight, Fog, MeshPhongMaterial, PointLight, SpotLight, TextureLoader } from "three";

// Gen random data
const N = 20;
const arcsData = [] as object[];

// custom globe material
const globeMaterial = new MeshPhongMaterial();
globeMaterial.bumpScale = 10;
globeMaterial.color = new Color("#160E32");
globeMaterial.emissive = new Color("#160E32");
globeMaterial.emissiveIntensity = 0.1;
globeMaterial.shininess = 0.7;

interface CustomGlobeProps {
    arcsData: Array<Record<string, any>>,
    setVisible?: () => void
}

export default function CustomGlobe({ arcsData, setVisible }: CustomGlobeProps) {

    const globeEl = useRef<any>();
    const [hex, setHex] = useState<any>({ features: [] });
    const [globeReady, setGlobeReady] = useState<boolean>(false);

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
        globe.controls().autoRotateSpeed = -0.4;
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
        camera.position.y = 100;

        globe.scene().add(camera);
        setGlobeReady(true);
    }, []);

    useEffect(() => {
        const globe = globeEl.current;
        globe.pointOfView({ lat: "39.099724", lng: "-94.578331" }, 5000)
    }, [globeReady])

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
            animateIn={true}
            // COUNTRIES
            hexPolygonsData={hex.features}
            hexPolygonResolution={useCallback(() => 3, [])} //values higher than 3 makes it buggy
            hexPolygonMargin={useCallback(() => 0.6, [])} // you can mess with this to see smaller or bigger dots
            hexPolygonColor={useCallback(() => "rgba(255, 255, 255, 1)", [])}
            hexPolygonCurvatureResolution={useCallback(() => 7, [])}
            //ARCS
            arcsData={arcsData}
            arcColor={() => '#efdefa'}
            arcDashLength={() => 0.7}
            arcDashGap={() => 0.7}
            arcDashInitialGap={1}
            arcDashAnimateTime={() => 2000}
            arcAltitudeAutoScale={() => 0.2}
            //arcAltitude={() => 0.05}
            arcCircularResolution={4}
            arcStroke={() => 0.7}
        />
    );
}
