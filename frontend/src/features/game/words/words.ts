// src/constants/word.ts หรือ "@/constants/word.ts"
export type WordDifficulty = "easy" | "medium" | "hard";

export interface WordEntry {
  word: string;
  difficulty: WordDifficulty;
}

export const getRandomWordByDifficulty = (
  difficulty: WordDifficulty
): WordEntry => {
  const filtered = wordPool.filter((entry) => entry.difficulty === difficulty);
  return filtered[Math.floor(Math.random() * filtered.length)];
};

export const wordPool: WordEntry[] = [
  // Easy 100 คำ
  { word: "cat", difficulty: "easy" },
  { word: "dog", difficulty: "easy" },
  { word: "sun", difficulty: "easy" },
  { word: "run", difficulty: "easy" },
  { word: "pen", difficulty: "easy" },
  { word: "box", difficulty: "easy" },
  { word: "hat", difficulty: "easy" },
  { word: "cup", difficulty: "easy" },
  { word: "car", difficulty: "easy" },
  { word: "bed", difficulty: "easy" },
  { word: "fish", difficulty: "easy" },
  { word: "tree", difficulty: "easy" },
  { word: "bird", difficulty: "easy" },
  { word: "milk", difficulty: "easy" },
  { word: "ball", difficulty: "easy" },
  { word: "shoe", difficulty: "easy" },
  { word: "door", difficulty: "easy" },
  { word: "cake", difficulty: "easy" },
  { word: "leaf", difficulty: "easy" },
  { word: "rain", difficulty: "easy" },
  { word: "star", difficulty: "easy" },
  { word: "frog", difficulty: "easy" },
  { word: "wind", difficulty: "easy" },
  { word: "fire", difficulty: "easy" },
  { word: "home", difficulty: "easy" },
  { word: "road", difficulty: "easy" },
  { word: "snow", difficulty: "easy" },
  { word: "book", difficulty: "easy" },
  { word: "coat", difficulty: "easy" },
  { word: "hand", difficulty: "easy" },
  { word: "moon", difficulty: "easy" },
  { word: "fox", difficulty: "easy" },
  { word: "map", difficulty: "easy" },
  { word: "key", difficulty: "easy" },
  { word: "egg", difficulty: "easy" },
  { word: "jam", difficulty: "easy" },
  { word: "bus", difficulty: "easy" },
  { word: "fan", difficulty: "easy" },
  { word: "bat", difficulty: "easy" },
  { word: "rat", difficulty: "easy" },
  { word: "toy", difficulty: "easy" },
  { word: "ice", difficulty: "easy" },
  { word: "owl", difficulty: "easy" },
  { word: "net", difficulty: "easy" },
  { word: "pen", difficulty: "easy" },
  { word: "lip", difficulty: "easy" },
  { word: "bag", difficulty: "easy" },
  { word: "hat", difficulty: "easy" },
  { word: "zip", difficulty: "easy" },
  { word: "bus", difficulty: "easy" },
  { word: "cap", difficulty: "easy" },
  { word: "mat", difficulty: "easy" },
  { word: "jet", difficulty: "easy" },
  { word: "pig", difficulty: "easy" },
  { word: "pot", difficulty: "easy" },
  { word: "pan", difficulty: "easy" },
  { word: "web", difficulty: "easy" },
  { word: "bug", difficulty: "easy" },
  { word: "jar", difficulty: "easy" },
  { word: "box", difficulty: "easy" },
  { word: "net", difficulty: "easy" },
  { word: "mud", difficulty: "easy" },
  { word: "sun", difficulty: "easy" },
  { word: "tip", difficulty: "easy" },
  { word: "van", difficulty: "easy" },
  { word: "lip", difficulty: "easy" },
  { word: "jet", difficulty: "easy" },
  { word: "cup", difficulty: "easy" },
  { word: "fan", difficulty: "easy" },
  { word: "rat", difficulty: "easy" },
  { word: "bat", difficulty: "easy" },
  { word: "log", difficulty: "easy" },
  { word: "bag", difficulty: "easy" },
  { word: "pit", difficulty: "easy" },
  { word: "cow", difficulty: "easy" },
  { word: "bee", difficulty: "easy" },
  { word: "web", difficulty: "easy" },
  { word: "owl", difficulty: "easy" },
  { word: "kid", difficulty: "easy" },
  { word: "cat", difficulty: "easy" },
  { word: "dog", difficulty: "easy" },
  { word: "sun", difficulty: "easy" },
  { word: "run", difficulty: "easy" },

  // Medium 100 คำ (ไม่มีซ้ำกับ easy)
  { word: "banana", difficulty: "medium" },
  { word: "pencil", difficulty: "medium" },
  { word: "rocket", difficulty: "medium" },
  { word: "jungle", difficulty: "medium" },
  { word: "planet", difficulty: "medium" },
  { word: "silver", difficulty: "medium" },
  { word: "danger", difficulty: "medium" },
  { word: "travel", difficulty: "medium" },
  { word: "hunter", difficulty: "medium" },
  { word: "forest", difficulty: "medium" },
  { word: "circle", difficulty: "medium" },
  { word: "window", difficulty: "medium" },
  { word: "garden", difficulty: "medium" },
  { word: "island", difficulty: "medium" },
  { word: "museum", difficulty: "medium" },
  { word: "letter", difficulty: "medium" },
  { word: "animal", difficulty: "medium" },
  { word: "market", difficulty: "medium" },
  { word: "family", difficulty: "medium" },
  { word: "office", difficulty: "medium" },
  { word: "season", difficulty: "medium" },
  { word: "summer", difficulty: "medium" },
  { word: "winter", difficulty: "medium" },
  { word: "autumn", difficulty: "medium" },
  { word: "school", difficulty: "medium" },
  { word: "player", difficulty: "medium" },
  { word: "stream", difficulty: "medium" },
  { word: "cloud", difficulty: "medium" },
  { word: "ocean", difficulty: "medium" },
  { word: "river", difficulty: "medium" },
  { word: "mountain", difficulty: "medium" },
  { word: "valley", difficulty: "medium" },
  { word: "desert", difficulty: "medium" },
  { word: "puzzle", difficulty: "medium" },
  { word: "battle", difficulty: "medium" },
  { word: "castle", difficulty: "medium" },
  { word: "kingdom", difficulty: "medium" },
  { word: "knight", difficulty: "medium" },
  { word: "library", difficulty: "medium" },
  { word: "mission", difficulty: "medium" },
  { word: "palace", difficulty: "medium" },
  { word: "pirate", difficulty: "medium" },
  { word: "quest", difficulty: "medium" },
  { word: "shield", difficulty: "medium" },
  { word: "soldier", difficulty: "medium" },
  { word: "temple", difficulty: "medium" },
  { word: "throne", difficulty: "medium" },
  { word: "tower", difficulty: "medium" },
  { word: "village", difficulty: "medium" },
  { word: "weapon", difficulty: "medium" },
  { word: "wizard", difficulty: "medium" },
  { word: "archer", difficulty: "medium" },
  { word: "armor", difficulty: "medium" },
  { word: "battlefield", difficulty: "medium" },
  { word: "champion", difficulty: "medium" },
  { word: "climate", difficulty: "medium" },
  { word: "commerce", difficulty: "medium" },
  { word: "culture", difficulty: "medium" },
  { word: "dungeon", difficulty: "medium" },
  { word: "economy", difficulty: "medium" },
  { word: "energy", difficulty: "medium" },
  { word: "enemy", difficulty: "medium" },
  { word: "fishing", difficulty: "medium" },
  { word: "fortress", difficulty: "medium" },
  { word: "hunting", difficulty: "medium" },
  { word: "journey", difficulty: "medium" },
  { word: "legend", difficulty: "medium" },
  { word: "magic", difficulty: "medium" },
  { word: "merchant", difficulty: "medium" },
  { word: "mining", difficulty: "medium" },
  { word: "monster", difficulty: "medium" },
  { word: "nature", difficulty: "medium" },
  { word: "oasis", difficulty: "medium" },
  { word: "paladin", difficulty: "medium" },
  { word: "pirates", difficulty: "medium" },
  { word: "plains", difficulty: "medium" },
  { word: "portal", difficulty: "medium" },
  { word: "priest", difficulty: "medium" },
  { word: "relic", difficulty: "medium" },
  { word: "rescue", difficulty: "medium" },
  { word: "riverbank", difficulty: "medium" },
  { word: "rogue", difficulty: "medium" },
  { word: "sorcerer", difficulty: "medium" },
  { word: "spell", difficulty: "medium" },
  { word: "strategy", difficulty: "medium" },
  { word: "treasure", difficulty: "medium" },
  { word: "tunnel", difficulty: "medium" },
  { word: "villager", difficulty: "medium" },
  { word: "warrior", difficulty: "medium" },
  { word: "wizardry", difficulty: "medium" },
  { word: "zephyr", difficulty: "medium" },

  // Hard 100 ประโยคสั้น (3-6 คำ)
  { word: "Time waits for no one", difficulty: "hard" },
  { word: "Knowledge is power", difficulty: "hard" },
  { word: "Honesty is the best policy", difficulty: "hard" },
  { word: "Actions speak louder than words", difficulty: "hard" },
  { word: "Practice makes perfect", difficulty: "hard" },
  { word: "Better late than never", difficulty: "hard" },
  { word: "Beauty is in the eye", difficulty: "hard" },
  { word: "Beggars can't be choosers", difficulty: "hard" },
  { word: "Every cloud has a silver lining", difficulty: "hard" },
  { word: "Don't count your chickens early", difficulty: "hard" },
  { word: "Easy come, easy go", difficulty: "hard" },
  { word: "Fortune favors the bold", difficulty: "hard" },
  { word: "Good things take time", difficulty: "hard" },
  { word: "Haste makes waste", difficulty: "hard" },
  { word: "Ignorance is bliss", difficulty: "hard" },
  { word: "It takes two to tango", difficulty: "hard" },
  { word: "Laughter is the best medicine", difficulty: "hard" },
  { word: "Look before you leap", difficulty: "hard" },
  { word: "No pain, no gain", difficulty: "hard" },
  { word: "One step at a time", difficulty: "hard" },
  { word: "Patience is a virtue", difficulty: "hard" },
  { word: "Practice what you preach", difficulty: "hard" },
  { word: "Rome wasn't built in a day", difficulty: "hard" },
  { word: "The early bird catches the worm", difficulty: "hard" },
  { word: "Two heads are better than one", difficulty: "hard" },
  { word: "When in Rome, do as Romans", difficulty: "hard" },
  { word: "You can't judge a book", difficulty: "hard" },
  { word: "Actions speak louder than words", difficulty: "hard" },
  { word: "Don't bite the hand that feeds", difficulty: "hard" },
  { word: "Don't put all eggs in one basket", difficulty: "hard" },
  { word: "Every dog has its day", difficulty: "hard" },
  { word: "Familiarity breeds contempt", difficulty: "hard" },
  { word: "If it ain't broke, don't fix it", difficulty: "hard" },
  { word: "It's no use crying over spilt milk", difficulty: "hard" },
  { word: "Keep your friends close", difficulty: "hard" },
  { word: "Knowledge is power", difficulty: "hard" },
  { word: "Let sleeping dogs lie", difficulty: "hard" },
  { word: "No news is good news", difficulty: "hard" },
  { word: "Old habits die hard", difficulty: "hard" },
  { word: "One man's trash is another man's treasure", difficulty: "hard" },
  { word: "Out of sight, out of mind", difficulty: "hard" },
  { word: "Practice makes perfect", difficulty: "hard" },
  { word: "Seeing is believing", difficulty: "hard" },
  { word: "The pen is mightier than the sword", difficulty: "hard" },
  { word: "There's no place like home", difficulty: "hard" },
  { word: "Time heals all wounds", difficulty: "hard" },
  { word: "To each his own", difficulty: "hard" },
  { word: "Too many cooks spoil the broth", difficulty: "hard" },
  { word: "Two wrongs don't make a right", difficulty: "hard" },
  { word: "Variety is the spice of life", difficulty: "hard" },
  { word: "What goes around comes around", difficulty: "hard" },
  { word: "Where there's smoke, there's fire", difficulty: "hard" },
  { word: "You can't have your cake", difficulty: "hard" },
  { word: "You reap what you sow", difficulty: "hard" },
  { word: "Actions speak louder than words", difficulty: "hard" },
  { word: "All good things must end", difficulty: "hard" },
  { word: "Barking up the wrong tree", difficulty: "hard" },
  { word: "Blood is thicker than water", difficulty: "hard" },
  { word: "Cross that bridge when you come to it", difficulty: "hard" },
  { word: "Curiosity killed the cat", difficulty: "hard" },
  { word: "Don't judge a book by its cover", difficulty: "hard" },
  { word: "Don't put off until tomorrow", difficulty: "hard" },
  { word: "Every cloud has a silver lining", difficulty: "hard" },
  { word: "Give someone the benefit of the doubt", difficulty: "hard" },
  { word: "If you can't beat them, join them", difficulty: "hard" },
  { word: "Ignorance is bliss", difficulty: "hard" },
  { word: "It takes two to tango", difficulty: "hard" },
  { word: "Jump on the bandwagon", difficulty: "hard" },
  { word: "Keep your eyes peeled", difficulty: "hard" },
  { word: "Let the cat out of the bag", difficulty: "hard" },
  { word: "Miss the boat", difficulty: "hard" },
  { word: "On thin ice", difficulty: "hard" },
  { word: "Once in a blue moon", difficulty: "hard" },
  { word: "Pull someone's leg", difficulty: "hard" },
  { word: "Put all your eggs in one basket", difficulty: "hard" },
  { word: "Raining cats and dogs", difficulty: "hard" },
  { word: "See eye to eye", difficulty: "hard" },
  { word: "Speak of the devil", difficulty: "hard" },
  { word: "The ball is in your court", difficulty: "hard" },
  { word: "The best of both worlds", difficulty: "hard" },
  { word: "The last straw", difficulty: "hard" },
  { word: "There's a method to the madness", difficulty: "hard" },
];
