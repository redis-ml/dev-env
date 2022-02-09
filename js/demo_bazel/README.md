# Awesome Babylon.js Starter

> An opinionated starter template for Babylon.js projects that require an HTML user interface and/or surrounding static web page.

## Conventions

This template follows the [same coding conventions](https://doc.babylonjs.com/how_to/approved_naming_conventions) stipulated for contributions to the Babylon.js project itself.


## Tooling

* [Babylon.js](https://www.babylonjs.com/)
* [TypeScript](https://www.typescriptlang.org/)
* [SCSS styles](https://sass-lang.com/)
* [PUG templates](https://pugjs.org/)
* [Node.js](https://nodejs.org/)
* [Webpack](https://webpack.js.org/)
* [Babel](https://babeljs.io/)

## Structure

* `dist/` - Minimised & optimised production build files for distribution.
* `docs/` - Documentation.
* `src/` - Source files.
    * `app/` - Application code.
        * `declarations.d.ts` - Ambient type declarations.
        * `main.ts` - Main application entry point.
        * `WebGL.ts` - Babylon.js sample scene.
    * `assets/` - Scene assets.
        * `fonts/`
        * `images/`
        * `meshes/`
        * `sounds/`
        * `textures/`
        * `videos/`
    * `data/` - Data in JSON format.
    * `styles/` - SCSS styles.
        * `imports/`
        * `index.scss`
    * `templates/` - PUG templates.
        * `imports/`
        * `index.pug`
    * `index.ts` - Import everything - code, templates & styles.

## Installation

```shell
npm install
```

## Development

```shell
npm start
```

## Production

```shell
npm run build
```

## Check Types

```shell
npm run type-check
```

## Babylon.js Resources

* [Babylon.js Website](https://www.babylonjs.com/)
* [Babylon.js Docs](https://doc.babylonjs.com/)
* [Babylon.js API Docs](https://doc.babylonjs.com/api/)
* [Babylon.js Forum](https://forum.babylonjs.com/)
* [Babylon.js Old Forum Archive](https://www.html5gamedevs.com/forum/16-babylonjs/)
* [Babylon.js Github Repo](https://github.com/BabylonJS/Babylon.js)
* [Babylon.js Playground](https://www.babylonjs-playground.com/)
* [Babylon.js Sandbox](https://sandbox.babylonjs.com/)
* [Babylon.js Node Material Editor](https://nme.babylonjs.com/)
* [Blender to Babylon.js Exporter](https://github.com/BabylonJS/BlenderExporter)
