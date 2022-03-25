const schedule = require("node-schedule")

const scheduleCronstyle = () => {
     schedule.scheduleJob('*/5 * * * * *',()=>{
        console.log('scheduleCronstyle:' + new Date());
    });
}

let ok = true;

scheduleCronstyle();

setTimeout(()=>{console.log('对不起, 要你久候' + new Date())}, 1000 )

// while (ok) {
//     setTimeout(()=>{console.log('对不起, 要你久候')}, 1000 )
//     console.log("432143214")
// }

setInterval(()=>{console.log('对不起, 要你久候')}, 1000 )

// console.log("dfsdfdsafds")


// for (let i= 0; i< 4;i++) {
//
//     setInterval(()=>{console.log('对不起, 要你久候' + i)}, 1000 )
// }
