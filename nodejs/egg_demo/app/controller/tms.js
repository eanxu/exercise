const Controller = require('egg').Controller;

class TmsController extends Controller {
    async xyz() {
        const ctx = this.ctx
        const x = ctx.params.x
        const y = ctx.params.y
        const z = ctx.params.z
        let png_buffer = await ctx.service.tms.xyz(x, y, z);
        ctx.body = png_buffer;
        ctx.status = 200;
        ctx.set('content-type', 'image/png');
        ctx.set('content-length', png_buffer.length);
    }
}

module.exports = TmsController
