export function createFooter() {
    
// Create the footer
var footer = document.createElement('footer');
footer.className = 'footer';

// Create the copyright paragraph
var p = document.createElement('p');
p.textContent = '© 2023 \'O1 FOUNDERS - FORUM\'. All rights reserved.';
footer.appendChild(p);

// Create the team links
var team = document.createElement('a');
team.textContent = 'Team';
footer.appendChild(team);

var teamMembers = ['spanchen', 'schana', 'nkorba', 'mkoseoglu', 'vkyrychu'];
for (var i = 0; i < teamMembers.length; i++) {
    var separator = document.createTextNode(' | ');
    footer.appendChild(separator);

    var member = document.createElement('a');
    member.textContent = teamMembers[i];
    footer.appendChild(member);
}

// Append the footer to the body
document.body.appendChild(footer);

}