# Mystic

> A small CLI game to improve my Golang skills.

## Description

This is a fantasy 1v1 game inspired by Meteorfall and DnD.
There are four classes of characters to choose from (Knight, Mage, Archer and Rogue).
Each of these characters have set health, energy, and attacks.
The purpose of the game is to defeat the other player.
Each turn the players can choose to attack, defend or rest.

- Attack - choose randomly one of the player's attacks. If the user has enough energy - attack, else nothing happens;
- Defend - assign a random defence amount to the player. If the attack doesn't use the entire defence amount - regenerate a small amount of energy;
- Rest - regenerate a big amount of energy.

## Classes

### Knight

#### Attacks:

- Kick
- Slash
- Dash
- Blitz

### Archer

#### Attacks:

- Snipe
- Electric arrow
- Triple shot
- Arrow shower
