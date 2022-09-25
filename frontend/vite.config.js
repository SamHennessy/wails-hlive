/** @type {import('vite').UserConfig} */
export default {
    build: {
        rollupOptions: {
            output: {
                entryFileNames: `assets/[name].js`,
                // TODO: turn off chunk? Or scan for them?
                //       Maybe just scan for all css and js in non-dev?
                chunkFileNames: `assets/[name].js`,
                assetFileNames: (assetInfo) => {
                    let extType = assetInfo.name.split('.').at(1);
                    if (/png|jpe?g|svg|gif|tiff|bmp|ico/i.test(extType)) {
                        extType = 'images';
                    }
                    if (/woff2?|eot|ttf/i.test(extType)) {
                        extType = 'fonts';
                    }
                    return `assets/${extType}/[name][extname]`;
                }
            }
        },
    }
}
