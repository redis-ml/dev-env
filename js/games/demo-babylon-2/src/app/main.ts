import WebGL from './WebGL'

const renderCanvasId = 'render-canvas';

window.addEventListener('DOMContentLoaded', () => {

    const renderCanvas = document.getElementById(renderCanvasId);
    if (renderCanvas) {
        const webGL = new WebGL(renderCanvas as HTMLCanvasElement);
        webGL.loadAssets();
    }
    else {
        console.error(`Render canvas id "${renderCanvasId}" not found`);
    }
});
