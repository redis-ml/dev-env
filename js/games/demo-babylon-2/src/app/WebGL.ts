import { Engine } from "@babylonjs/core/Engines/engine";
import { Scene } from "@babylonjs/core/scene";
import { Vector3, Color3 } from "@babylonjs/core/Maths/math";
import { FreeCamera } from "@babylonjs/core/Cameras/freeCamera";
import { HemisphericLight } from "@babylonjs/core/Lights/hemisphericLight";
import { MeshBuilder } from "@babylonjs/core/Meshes/meshBuilder";
import { StandardMaterial } from "@babylonjs/core/Materials/standardMaterial"

export default class WebGL {

    private _engine: Engine;
    private _scene: Scene;

    /**
     * Creates an instance of WebGL.
     * 
     * @param {HTMLCanvasElement} renderCanvas
     * @memberof WebGL
     */
    constructor(renderCanvas: HTMLCanvasElement) {

        // Instantiate engine and scene.
        const options = {
            // See: https://doc.babylonjs.com/api/interfaces/babylon.engineoptions
            disableWebGL2Support: false, // Set true for webGL1 fallback testing.
            useHighPrecisionFloats: true // Set false for older device testing.
        }
        this._engine = new Engine(renderCanvas, true, options);
        this._scene = new Scene(this._engine);

        // Enable Babylon.js debug inspector for dev only (see webpack.config.dev.js).
        /* babylonjs-inspector */
    }

    /**
     * Load scene assets.
     *
     * @memberof WebGL
     */
    public loadAssets(): void {
        //let assetsLoader = new Loader(this._scene, this._assets);
        //assetsLoader.load(() => {
            // On loading finished ...
            this.createScene();
            this.renderScene();
        //});
    }

    /**
     * Create scene from loaded scene assets.
     *
     * @memberof WebGL
     */
    public createScene(): void {

        // Create camera.
        let camera = new FreeCamera('camera', new Vector3(0, 5,-10), this._scene);
        camera.setTarget(Vector3.Zero());
        camera.attachControl(this._engine.getRenderingCanvas() as HTMLElement, false);

        // Create light.
        new HemisphericLight('light', new Vector3(0,1,0), this._scene);

        // Create ground.
        let ground = MeshBuilder.CreateGround('ground', {width: 6, height: 6, subdivisions: 2}, this._scene);

        // Create sphere.
        let sphere = MeshBuilder.CreateSphere('sphere', {segments: 16, diameter: 2}, this._scene);
        sphere.position.y = 1;

        // Set material.
        let material = new StandardMaterial("material", this._scene);
        material.diffuseColor = new Color3(1, 1, 1);
        sphere.material = material;
        ground.material = material;

        // Animate sphere.
        let alpha = 0;
        this._scene.onBeforeRenderObservable.add(() => {
            sphere.scaling.y = Math.cos(alpha);
            alpha += 0.01;
        });
    }

    /**
     * Render the created scene.
     *
     * @memberof WebGL
     */
    public renderScene(): void {

        // Start render loop.
        this._engine.runRenderLoop(() => {
            this._scene.render();
        });

        // Add resize event listener.
        window.addEventListener('resize', () => {
            this._engine.resize();
        });
    }
}
