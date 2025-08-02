import { create } from "zustand"

interface Critique {
  review: string
  scores: {
    originality: number
    scalability: number
    feasibility: number
  }
}

interface AppState {
  idea: string
  critique: Critique | null
  isLoading: boolean
  isDefending: boolean
  isImproving: boolean
  setIdea: (idea: string) => void
  submitIdea: () => Promise<void>
  defendIdea: () => Promise<void>
  improveIdea: () => Promise<void>
  reset: () => void
}

// Mock API calls - replace with actual API endpoints
const mockSubmitIdea = async (idea: string): Promise<Critique> => {
  await new Promise((resolve) => setTimeout(resolve, 2000))
  return {
    review: `Your idea "${idea.slice(0, 50)}..." shows some promise but has significant flaws. The concept lacks differentiation in a crowded market and faces substantial execution challenges. The target audience isn't clearly defined, and the monetization strategy appears weak. While the core problem you're addressing is valid, your proposed solution is overly simplistic and doesn't account for real-world complexities. You'll need to dig deeper into user research and competitive analysis before this becomes viable.`,
    scores: {
      originality: Math.floor(Math.random() * 40) + 30,
      scalability: Math.floor(Math.random() * 50) + 25,
      feasibility: Math.floor(Math.random() * 60) + 20,
    },
  }
}

const mockDefendIdea = async (): Promise<string> => {
  await new Promise((resolve) => setTimeout(resolve, 1500))
  return "Actually, upon further analysis, this idea has several underappreciated strengths. The timing might be perfect given current market trends, and the simplicity could be a feature rather than a bug. Many successful products started with seemingly basic concepts that solved real problems elegantly."
}

const mockImproveIdea = async (): Promise<string> => {
  await new Promise((resolve) => setTimeout(resolve, 1500))
  return "Here's how to make this idea stronger: 1) Focus on a specific niche market first, 2) Add a freemium model with premium features, 3) Implement social proof mechanisms, 4) Create a waitlist to build anticipation, 5) Partner with established players for distribution. Consider pivoting to a B2B model if B2C proves challenging."
}

export const useAppStore = create<AppState>((set, get) => ({
  idea: "",
  critique: null,
  isLoading: false,
  isDefending: false,
  isImproving: false,

  setIdea: (idea) => set({ idea }),

  submitIdea: async () => {
    const { idea } = get()
    if (!idea.trim()) return

    set({ isLoading: true })
    try {
      const critique = await mockSubmitIdea(idea)
      set({ critique, isLoading: false })
    } catch (error) {
      set({ isLoading: false })
    }
  },

  defendIdea: async () => {
    set({ isDefending: true })
    try {
      const defense = await mockDefendIdea()
      const { critique } = get()
      if (critique) {
        set({
          critique: { ...critique, review: defense },
          isDefending: false,
        })
      }
    } catch (error) {
      set({ isDefending: false })
    }
  },

  improveIdea: async () => {
    set({ isImproving: true })
    try {
      const improvement = await mockImproveIdea()
      const { critique } = get()
      if (critique) {
        set({
          critique: { ...critique, review: improvement },
          isImproving: false,
        })
      }
    } catch (error) {
      set({ isImproving: false })
    }
  },

  reset: () =>
    set({
      idea: "",
      critique: null,
      isLoading: false,
      isDefending: false,
      isImproving: false,
    }),
}))
