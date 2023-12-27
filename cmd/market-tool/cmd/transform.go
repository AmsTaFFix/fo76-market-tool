package cmd

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// transformCmd represents the transform command
var transformCmd = &cobra.Command{
	Use:   "transform",
	Short: "transforms current json output to feed it too fo76market web-site to make it work again",
	RunE:  transformFunc,
	Args:  cobra.ExactArgs(2),
}

func init() {
	rootCmd.AddCommand(transformCmd)
}

var replaces = map[string]string{
	"Damage increases up to +24% as you fill your hunger and thirst meters": "Damage increases as you fill your hunger and thirst meters",
	"+25% Weapon Speed": "25% faster fire rate",
	"Damage Increases up to +95% as Health Decreases": "Damage increases as health decreases",
	"+15% Reload Speed":                                                          "15% faster reload",
	"+50% Chance to Hit a Target in V.A.T.S.":                                    "+50% V.A.T.S. hit chance",
	"Damage Increases per Addiction up to +50% Currently +0%":                    "Damage increases when suffering from addictions",
	"+100% Faster Movement Speed While Aiming":                                   "Faster movement speed while aiming",
	"-25% Action Point Cost":                                                     "25% less V.A.T.S. Action Point cost",
	"+50% Armor Penetration":                                                     "Ignores 50% of your target's armor",
	"+50% Critical Damage":                                                       "V.A.T.S. critical hits do +50% damage",
	"Bullets Explode for 20% Weapon Damage":                                      "Bullets explode for area damage",
	"Bullets Explode for 3% Weapon Damage":                                       "Bullets explode for area damage",
	"Damage Increases up to +25% as you gain Mutations":                          "Damage increased by 5% for each mutation",
	"Replenish 15 Action Points with Each Kill":                                  "Replenish Action Points with each kill",
	"The last round in a magazine has a 25% chance to Deal +100% Damage":         "The last round in a magazine has a 25% chance to deal 2x DMG",
	"10% Chance to Generate a Stealth Field for 2 seconds when Hitting a Target": "Hits have a chance to generate a Stealth Field",
	"If Not In Combat, +100% V.A.T.S. Accuracy at +50% Action Points Cost":       "If not in combat, +100% V.A.T.S. accuracy at +50% AP cost",
	"+15 Bonus V.A.T.S. Critical Charge":                                         "Your V.A.T.S. critical meter fills 15% faster",
	"+50% Damage to Insects +50% Damage to Mirelurks":                            "+50% damage to Mirelurks and bugs",
	"-90% Weight":         "90% reduced weight",
	"+300% Ammo Capacity": "Quadruple ammo capacity",
	"Up to +50% Damage based on Caps you have":                        "Damage increases as caps increases",
	"Restore 2% Health over 2 seconds when you Hit a Target":          "Gain brief health regeneration when you hit an enemy",
	"+50% Damage at Night":                                            "Damage increases with the night",
	"Reduce Your Target's Damage Output by 25% for 5 seconds":         "Reduce your target's damage output by 25% for 5s",
	"Damage Increases up to +25% as Health Increases":                 "Damage increases as health increases",
	"+1 Projectiles +25% Damage -150% Hip-Fire Accuracy +100% Recoil": "Shoots an additional projectile",
	"+50% Bash Damage": "Bashing damage increased by 50%",
	"+5% Damage after Each Consecutive Hit on the Same Target, up to +45%": "Damage increased after each consecutive hit on the same target",
	"+100% Damage Against Targets with Full Health":                        "Double damage if target is full health",
	"Damage Increases up to +50% as Damage Resistance Decreases":           "Lower Damage Resistance increases damage dealt",
	"V.A.T.S. Criticals will Heal You and Your Group by 5% Health":         "V.A.T.S. crits will heal you and your group",
	"+40% Weapon Speed":                                                                              "40% faster swing speed",
	"+40% Power Attack Damage":                                                                       "40% more power attack damage",
	"+50% Melee Damage Reflection While Blocking":                                                    "Reflects 50% of melee damage back while blocking",
	"-40% Damage Taken While Power Attacking":                                                        "Take 40% less damage while power attacking",
	"Weapons break 50% slower":                                                                       "Breaks 50% slower",
	"-15% Damage Taken While Blocking":                                                               "Take 15% less damage while blocking",
	"+25% Melee Damage While Not Moving":                                                             "+25% damage while standing still",
	"Gain up to +3 to All S.P.E.C.I.A.L. Stats (except END) when Health is Low":                      "Gain up to +3 to all stats (except END) when low health",
	"Grants up to +35 Energy Resistance and Damage Resistance, the higher Your Health":               "Grants up to +35 Energy Resistance and Damage Resistance, the higher your health",
	"+25% Reduced Disease Chance From Environmental Hazards":                                         "+25% Environmental Disease Resistance",
	"Grants up to +35 Energy Resistance and Damage Resistance, the Lower Your Health":                "Grants up to +35 Energy Resistance and Damage Resistance, the lower your health",
	"Being Hit in Melee Generates a Stealth Field Once Every 30 Seconds":                             "Being hit in melee generates a Stealth Field once per 30 seconds",
	"+5% Effectiveness of Stimpaks, RadAway, and Rad-X":                                              "Stimpaks, RadAway, and Rad-X are 5% more effective",
	"When Incapacitated, Gain a 50% Chance to Revive Yourself with a Stimpak, Once Every 60 seconds": "When incapacitated, gain a 50% chance to revive yourself with a Stimpak, once every minute",
	"5% chance to deal 12 Energy damage per second for 4 seconds to Melee Attackers.":                "5% chance to deal 100 Energy DMG to melee attackers",
	"Increases Damage Reduction up to +6% as you Fill Your Hunger and Thirst Meters":                 "Increases Damage Reduction up to 6% as you fill your hunger and thirst meters",
	"+25% Less Noise While Sneaking +25% Reduce Detection Chance":                                    "Become harder to detect while sneaking",
	"-15% Damage from Mirelurks and Insects":                                                         "-15% damage from Mirelurks and bugs",
	"+40 Damage Resistance and Energy Resistance at Night":                                           "Damage and Energy Resistance increase at night",
	"-50% Fall Damage":                 "Reduces falling damage by 50%",
	"Reduced weapon weight":            "Weapon weights reduced by 20%",
	"+0.25% Radiation Damage Recovery": "Slowly regen radiation damage while not in combat",
	"5% chance to deal 12 Fire damage per second for 4 seconds to Melee Attackers.": "5% chance to deal 100 Fire DMG to melee attackers",
	"+0.5% Heal Rate": "Slowly regenerate health while not in combat",
	"Become Invisible While Sneaking and Not Moving":                                  "Blend with the environment while sneaking and not moving",
	"+10 Damage Resistance and Energy Resistance While Mutated":                       "+10 Damage Resistance and Energy Resistance if you are mutated",
	"Hunger and Thirst Grow 10% Slower":                                               "UNKNOWN",
	"5% chance to deal 12 Poison damage per second for 4 seconds to Melee Attackers.": "5% chance to deal 100 Poison DMG to melee attackers",
	"+5% Action Point Regen":                                                          "Increases Action Point refresh speed",
	"75% Chance to Reduce Damage by 15% While Not Moving":                             "75% chance to reduce damage by 15% while standing still",
	"Grants up to +20 Energy Resistance and Damage Resistance, the higher Your Caps":  "Grants up to +20 Energy Resistance and Damage Resistance, the higher your caps",
	"5% chance to deal 12 Cryo damage per second for 4 seconds to Melee Attackers.":   "5% chance to deal 100 Frost DMG to melee attackers",
	"Increases Size of Sweet-Spot While Picking Locks by 2":                           "Increases size of sweet-spot while picking locks",
}

func transformFunc(cmd *cobra.Command, args []string) error {
	filePath := args[0]
	f, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("can't open file '%s': %w", filePath, err)
	}
	defer f.Close()

	obj := AutoGenerated{}
	jsonDecoder := json.NewDecoder(f)
	err = jsonDecoder.Decode(&obj)
	if err != nil {
		return fmt.Errorf("can't parse json content: %w", err)
	}

	for _, character := range obj.CharacterInventories {
		for _, items := range [][]*Item{character.PlayerInventory, character.StashInventory} {
			for _, item := range items {
				if item.ContainerID > 2147483647 {
					item.ContainerID = 2147483647
				}
				if item.QuickSwapAmmo > 2147483647 {
					item.QuickSwapAmmo = 2147483647
				}
				if item.MaximumHealth > 2147483647 {
					item.MaximumHealth = 2147483647
				}
				if item.IsLegendary {
					for _, cardEntry := range item.ItemCardEntries {
						if cardEntry.Text == "DESC" {
							parts := make([]string, 0, 3)
							for _, line := range strings.Split(cardEntry.Value, "\n") {
								if strings.Contains(line, "¬") {
									parts = append(parts, strings.ReplaceAll(line, "¬", ""))
								}
							}

							for partIdx, part := range parts {
								for newText, oldText := range replaces {
									if strings.Contains(strings.ToLower(part), strings.ToLower(newText)) {
										if item.FilterFlag == 1<<3 { // Armor
											if strings.Contains(strings.ToLower(part), strings.ToLower("-15% Damage Taken While Blocking")) {
												oldText = "Reduces damage while blocking by 15%"
											}
											if strings.Contains(strings.ToLower(part), strings.ToLower("-90% Weight")) {
												oldText = "Weighs 90% less and does not count as armor for the Chameleon mutation"
											}
											if oldText == "UNKNOWN" {
												oldText = "+25 Radiation Resistance"
												fmt.Printf("item '%s', changed star from '%s' to '%s' because system do not know about new one\n",
													item.Text, newText, oldText)
											}
										}
										parts[partIdx] = oldText
									}
								}
							}

							value := strings.Join(parts, "\n")
							cardEntry.Value = value
						}
					}
				}
				if !item.IsLegendary {
					for _, cardEntry := range item.ItemCardEntries {
						if cardEntry.Text == "DESC" {
							cardEntry.Value = ""
						}
					}
				}
			}
		}
	}

	fDest, err := os.Create(args[1])
	if err != nil {
		return fmt.Errorf("can't open/create file '%s' for result: %w", args[1], err)
	}
	defer fDest.Close()
	jsonEncoder := json.NewEncoder(fDest)
	err = jsonEncoder.Encode(obj)
	if err != nil {
		return fmt.Errorf("can't encode result: %w", err)
	}

	return nil
}

type Item struct {
	Text                    string  `json:"text"`
	ServerHandleID          int64   `json:"serverHandleId"`
	Count                   int     `json:"count"`
	ItemValue               int     `json:"itemValue"`
	EquipState              int     `json:"equipState"`
	IsWeightless            bool    `json:"isWeightless"`
	FilterFlag              int     `json:"filterFlag"`
	IsNew                   bool    `json:"isNew"`
	CurrentHealth           float64 `json:"currentHealth"`
	MaximumHealth           int     `json:"maximumHealth"`
	Durability              int     `json:"durability"`
	ItemLevel               int     `json:"itemLevel"`
	QuickSwapAmmo           int     `json:"quickSwapAmmo"`
	CanFavorite             bool    `json:"canFavorite"`
	Favorite                bool    `json:"favorite"`
	IsLegendary             bool    `json:"isLegendary"`
	NumLegendaryStars       int     `json:"numLegendaryStars"`
	TaggedForSearch         bool    `json:"taggedForSearch"`
	ScrapAllowed            bool    `json:"scrapAllowed"`
	IsAutoScrappable        bool    `json:"isAutoScrappable"`
	CanGoIntoScrapStash     bool    `json:"canGoIntoScrapStash"`
	QuickSwap               bool    `json:"quickSwap"`
	IsTradable              bool    `json:"isTradable"`
	SingleItemTransfer      bool    `json:"singleItemTransfer"`
	IsQuestItem             bool    `json:"isQuestItem"`
	IsSharedQuestItem       bool    `json:"isSharedQuestItem"`
	Weight                  float64 `json:"weight"`
	WeaponDisplayAccuracy   float64 `json:"weaponDisplayAccuracy"`
	WeaponDisplayRange      float64 `json:"weaponDisplayRange"`
	WeaponDisplayRateOfFire float64 `json:"weaponDisplayRateOfFire"`
	Damage                  float64 `json:"damage"`
	IsCurrency              bool    `json:"isCurrency"`
	IsOffered               bool    `json:"isOffered"`
	OfferValue              int     `json:"offerValue"`
	IsRequested             bool    `json:"isRequested"`
	DeclineReason           int     `json:"declineReason"`
	IsSpoiled               bool    `json:"isSpoiled"`
	IsLearnedRecipe         bool    `json:"isLearnedRecipe"`
	IsSetItem               bool    `json:"isSetItem"`
	IsSetBonusActive        bool    `json:"isSetBonusActive"`
	ContainerID             int64   `json:"containerID"`
	ItemCategory            int     `json:"itemCategory"`
	VendingData             struct {
		MachineType              int  `json:"machineType"`
		Price                    int  `json:"price"`
		IsVendedOnOtherMachine   bool `json:"isVendedOnOtherMachine"`
		IsUsedOnOtherCampMachine bool `json:"isUsedOnOtherCampMachine"`
		AssignEnabled            bool `json:"assignEnabled"`
		UnassignEnabled          bool `json:"unassignEnabled"`
		ReadOnly                 bool `json:"readOnly"`
	} `json:"vendingData"`
	Rarity          int `json:"rarity"`
	ItemCardEntries []*struct {
		Text              string        `json:"text"`
		Value             string        `json:"value"`
		DamageType        int           `json:"damageType"`
		Difference        float64       `json:"difference"`
		DiffRating        int           `json:"diffRating"`
		Precision         int           `json:"precision"`
		ShowAsDescription bool          `json:"showAsDescription"`
		Duration          int           `json:"duration"`
		ProjectileCount   int           `json:"projectileCount"`
		Components        []interface{} `json:"components"`
	} `json:"ItemCardEntries"`
}
type AutoGenerated struct {
	CharacterInventories map[string]struct {
		CharacterInfoData struct {
			Level int    `json:"level"`
			Name  string `json:"name"`
		} `json:"CharacterInfoData"`
		AccountInfoData struct {
			Name string `json:"name"`
		} `json:"AccountInfoData"`
		StashInventory  []*Item `json:"stashInventory"`
		PlayerInventory []*Item `json:"playerInventory"`
	} `json:"characterInventories"`
	Version float64 `json:"version"`
	ModName string  `json:"modName"`
}
