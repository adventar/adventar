module.exports = {
  moduleNameMapper: {
    "^~/(.*)$": "<rootDir>/$1"
  },
  moduleFileExtensions: ["js", "ts", "vue", "json"],
  transform: {
    "^.+\\.tsx?$": "ts-jest",
    "^.+\\.vue$": "vue-jest"
  }
};
