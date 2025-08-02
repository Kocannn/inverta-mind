"use client"

import type React from "react"
import { Brain, Loader2, Shield, Lightbulb, RotateCcw } from "lucide-react"
import { Button } from "@/components/ui/button"
import { Textarea } from "@/components/ui/textarea"
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card"
import { ScoreCard } from "@/components/score-card"
import { ThemeToggle } from "@/components/theme-toggle"
import { useAppStore } from "@/lib/store"

export default function HomePage() {
  const { idea, critique, isLoading, isDefending, isImproving, setIdea, submitIdea, defendIdea, improveIdea, reset } =
    useAppStore()

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault()
    await submitIdea()
  }

  return (
    <div className="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 dark:from-gray-900 dark:to-gray-800">
      {/* Header */}
      <header className="border-b border-gray-200 dark:border-gray-700 bg-white/80 dark:bg-gray-900/80 backdrop-blur-sm sticky top-0 z-10">
        <div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            <div className="flex items-center space-x-3">
              <div className="p-2 bg-red-100 dark:bg-red-900/30 rounded-lg">
                <Brain className="h-6 w-6 text-red-600 dark:text-red-400" />
              </div>
              <div>
                <h1 className="text-xl font-bold text-gray-900 dark:text-white">Self-Debunking AI</h1>
                <p className="text-xs text-gray-500 dark:text-gray-400">Brutal honesty for better ideas</p>
              </div>
            </div>
            <div className="flex items-center space-x-2">
              {critique && (
                <Button variant="outline" size="sm" onClick={reset} className="hidden sm:flex bg-transparent">
                  <RotateCcw className="h-4 w-4 mr-2" />
                  New Idea
                </Button>
              )}
              <ThemeToggle />
            </div>
          </div>
        </div>
      </header>

      <main className="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8 py-8">
        {!critique ? (
          /* Input Section */
          <div className="max-w-2xl mx-auto">
            <div className="text-center mb-8">
              <h2 className="text-3xl font-bold text-gray-900 dark:text-white mb-4">Ready for some brutal honesty?</h2>
              <p className="text-lg text-gray-600 dark:text-gray-300">
                Submit your idea and get an unfiltered critique from our AI. No sugar-coating, just honest feedback to
                help you improve.
              </p>
            </div>

            <Card className="shadow-xl border-0 bg-white/70 dark:bg-gray-800/70 backdrop-blur-sm">
              <CardContent className="p-8">
                <form onSubmit={handleSubmit} className="space-y-6">
                  <div>
                    <label htmlFor="idea" className="block text-sm font-medium text-gray-700 dark:text-gray-300 mb-3">
                      Describe your idea
                    </label>
                    <Textarea
                      id="idea"
                      value={idea}
                      onChange={(e) => setIdea(e.target.value)}
                      placeholder="Tell us about your app, business, or project idea. Be as detailed as you want - the more context you provide, the better feedback you'll get..."
                      className="min-h-[200px] text-base resize-none border-gray-300 dark:border-gray-600 focus:border-red-500 dark:focus:border-red-400 focus:ring-red-500 dark:focus:ring-red-400"
                      disabled={isLoading}
                    />
                  </div>

                  <Button
                    type="submit"
                    disabled={!idea.trim() || isLoading}
                    className="w-full h-12 text-base font-medium bg-red-600 hover:bg-red-700 dark:bg-red-700 dark:hover:bg-red-600"
                  >
                    {isLoading ? (
                      <>
                        <Loader2 className="h-5 w-5 mr-2 animate-spin" />
                        Analyzing your idea...
                      </>
                    ) : (
                      "Get Brutal Feedback"
                    )}
                  </Button>
                </form>
              </CardContent>
            </Card>
          </div>
        ) : (
          /* Critique Section */
          <div className="space-y-8">
            {/* Mobile Reset Button */}
            <div className="sm:hidden">
              <Button variant="outline" size="sm" onClick={reset} className="w-full bg-transparent">
                <RotateCcw className="h-4 w-4 mr-2" />
                Submit New Idea
              </Button>
            </div>

            {/* Critique Card */}
            <Card className="shadow-xl border-0 bg-white/70 dark:bg-gray-800/70 backdrop-blur-sm">
              <CardHeader>
                <CardTitle className="text-xl font-bold text-gray-900 dark:text-white">AI Critique</CardTitle>
              </CardHeader>
              <CardContent className="space-y-6">
                <div className="prose prose-gray dark:prose-invert max-w-none">
                  <p className="text-gray-700 dark:text-gray-300 leading-relaxed">{critique.review}</p>
                </div>

                {/* Scores */}
                <div className="grid grid-cols-1 md:grid-cols-3 gap-6 pt-6 border-t border-gray-200 dark:border-gray-700">
                  <ScoreCard label="Originality" score={critique.scores.originality} />
                  <ScoreCard label="Scalability" score={critique.scores.scalability} />
                  <ScoreCard label="Feasibility" score={critique.scores.feasibility} />
                </div>

                {/* Action Buttons */}
                <div className="flex flex-col sm:flex-row gap-4 pt-6 border-t border-gray-200 dark:border-gray-700">
                  <Button
                    onClick={defendIdea}
                    disabled={isDefending || isImproving}
                    variant="outline"
                    className="flex-1 h-12 border-green-300 text-green-700 hover:bg-green-50 dark:border-green-600 dark:text-green-400 dark:hover:bg-green-900/20 bg-transparent"
                  >
                    {isDefending ? (
                      <>
                        <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                        Defending...
                      </>
                    ) : (
                      <>
                        <Shield className="h-4 w-4 mr-2" />
                        Defend this idea
                      </>
                    )}
                  </Button>

                  <Button
                    onClick={improveIdea}
                    disabled={isDefending || isImproving}
                    variant="outline"
                    className="flex-1 h-12 border-blue-300 text-blue-700 hover:bg-blue-50 dark:border-blue-600 dark:text-blue-400 dark:hover:bg-blue-900/20 bg-transparent"
                  >
                    {isImproving ? (
                      <>
                        <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                        Improving...
                      </>
                    ) : (
                      <>
                        <Lightbulb className="h-4 w-4 mr-2" />
                        Make it better
                      </>
                    )}
                  </Button>
                </div>
              </CardContent>
            </Card>
          </div>
        )}
      </main>
    </div>
  )
}
