import React from 'react'
interface ScoreCardProps {
  label: string
  score: number
}

export function ScoreCard({ label, score }: ScoreCardProps) {
  const getScoreColor = (score: number) => {
    if (score >= 70) return "text-green-600 dark:text-green-400"
    if (score >= 40) return "text-yellow-600 dark:text-yellow-400"
    return "text-red-600 dark:text-red-400"
  }

  const getBarColor = (score: number) => {
    if (score >= 70) return "bg-green-500"
    if (score >= 40) return "bg-yellow-500"
    return "bg-red-500"
  }

  return (
    <div className="space-y-2">
      <div className="flex justify-between items-center">
        <span className="text-sm font-medium text-gray-700 dark:text-gray-300">{label}</span>
        <span className={`text-lg font-bold ${getScoreColor(score)}`}>{score}</span>
      </div>
      <div className="w-full bg-gray-200 dark:bg-gray-700 rounded-full h-2">
        <div
          className={`h-2 rounded-full transition-all duration-500 ${getBarColor(score)}`}
          style={{ width: `${score}%` }}
        />
      </div>
    </div>
  )
}
