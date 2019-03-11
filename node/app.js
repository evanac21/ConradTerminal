const rls = require('readline-sync');
const ch = require('chalk');
const fs = require('fs');
const bcrypt = require('bcrypt');
var keypress = require('keypress')
log = console.log

runtime()
keypress(process.stdin);

// listen for the "keypress" event

function runtime() {
    log(ch.red("Conrad Terminal"))
    log(ch.blue("F7. Login"))
    log(ch.blue("F8. Register"))
    log(ch.blue("F9. Settings"))
    log(ch.blue("F10. Exit"))
    process.stdin.on('keypress', function(ch, key) {

        switch (key.name) {
            case "f7":
                var cp = checkPass()
                if (cp = true) {
                    fmain()
                } else {
                    console.log("Deauthed")
                    process.exit(0)
                }
                break;
            case "f8":
                console.log("Currently doesnt work!")
                process.stdin.pause()
                process.exit(0)
                break;
            case "f9":
                console.log("f9")
                break;
            case "f10":
                process.stdin.pause()
                process.exit(0)
                break;
        }
        if (key && key.ctrl && key.name == 'c') {
            process.stdin.pause();
        }
    });

    process.stdin.setRawMode(true);
    process.stdin.resume();

}

function checkPass() {
    var hash = fs.readFileSync("hash.txt")
    var password = rls.question("Password?: ")
    bcrypt.compare(password, hash, function(err, res) {
        return res;
    });
}

/*plainPassword = "bobcat"
var saltRounds = 10
bcrypt.hash(plainPassword, saltRounds, function(err, hash) {
    fs.writeFileSync("hash.txt", hash)
});
*/
function cls() {
    var lines = process.stdout.getWindowSize()[1];
    for (var i = 0; i < lines; i++) {
        console.log('\r\n');
    }
}

function fmain() {
    cls()
    console.log("Craig Morgan is my literal father")
}