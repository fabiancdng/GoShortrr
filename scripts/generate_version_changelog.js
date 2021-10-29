/**
 * Grabs only a the part of CHANGELOG.md that's relevant
 * for the current release and writes it to a (temporary)
 * own MD file that can than be attached to the release.
 */
const fs = require('fs');

console.log('Trying to read version-specific changelog from CHANGELOG.md...');
var changelog = fs.readFileSync('../CHANGELOG.md').toString().split('\n');

var versionLog = '';

var read = false;

for (lineIndex in changelog) {
    var lineContent = changelog[lineIndex];

    if (lineContent.startsWith('## ')) read = true;
    else if (lineContent.startsWith('<a name=') && lineIndex >= 5) break;

    if (read) versionLog += lineContent + '\n';
}

console.log('Writing version log to own file...');

fs.writeFileSync('../VERSION_LOG.md', versionLog);

console.log('Done!');