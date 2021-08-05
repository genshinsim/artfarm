# artfarm

This is a simple tool that calculates on average how many artifact drops it will take to farm a 4 piece set with the desired main stat and desired amount of certain sub stats.

Simply modify the config file and run the tool. Sample config as follows:

```json
{
    "main_stat": {
        "flower": "hp",
        "feather": "atk",
        "sand": "atk%",
        "goblet": "pyro%",
        "circlet": "cr"
    },
    "desired_subs": {
        "cr": 0.24,
        "cd": 0.3
    },
    "iterations": 100000,
    "workers": 24
}
```

The following is a list of acceptable input (**CASE SENSITIVE**) for stat types

```
def%
def
hp
hp%
atk
atk%
er
em
cr
cd
heal
pyro%
hydro%
cryo%
electro%
anemo%
geo%
phys%
```

Also note that the tool will not validate your input (i.e. feather with atk%). It will simply never reach a result since you can never acquire a feather with atk% main stat.

## Algorithm

The tool follows a very simple greedy algorithm where it will generate a random artifact, and keep it if the following criteria are met (replacing any existing):

- It is either on set (50/50 chance) OR if not on set, then it is a goblet piece with the desired main stat
- it has the desired main stat
- The total substat with this new random artifact is closer percentage wise to the desired substat than the existing artifact it is replacing.

Note in practice this will most likely overestimate the total drops required to reach the desired sub threshold.
