import { apiClient } from "./api"
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
  submitStreamIdea: () => Promise<void>
  postIdea: () => Promise<void>
  defendIdea: () => Promise<void>
  improveIdea: () => Promise<void>
  reset: () => void
  streamingContent: string
}

export const useAppStore = create<AppState>((set, get) => ({
  idea: "",
  critique: null,
  isLoading: false,
  isDefending: false,
  isImproving: false,
  streamingContent: "",

  setIdea: (idea) => set({ idea }),

  submitIdea: async () => {
    const { idea } = get();
    if (!idea.trim()) return;

    set({ isLoading: true });
    try {
      // Call the API to submit the idea - now returns properly formatted critique
      const critique = await apiClient.submitIdea(idea);

      // Set the critique directly since it's already properly formatted
      set({ critique, isLoading: false });
    } catch (error) {
      console.error("Error submitting idea:", error);
      set({ isLoading: false });
    }
  },


  postIdea: async () => {
    const { idea } = get();
    if (!idea.trim()) return;

    set({ isLoading: true, streamingContent: "" });
    try {
      // First post the idea to get an ID
      const response = await apiClient.PostIdea(idea);
      console.log("Idea posted successfully:", response);

      // Now start streaming with the received ID
      if (response && response.data && response.data.id) {
        console.log("Starting stream with ID:", response.data.id);

        apiClient.streamSubmitIdea(
          // Update UI langsung untuk setiap chunk
          (chunk) => {
            // Tambahkan efek mengetik dengan mengganti seluruh konten
            set({ streamingContent: chunk });

            // Auto-scroll ke bagian bawah jika di mobile/viewport sempit
            if (window.innerWidth < 768) {
              setTimeout(() => {
                window.scrollTo({
                  top: document.body.scrollHeight,
                  behavior: 'smooth'
                });
              }, 100);
            }
          },
          // Setelah selesai, simpan hasil lengkap
          (critique) => {
            set({
              critique,
              isLoading: false,
              streamingContent: ""
            });
          },
          response.data.id
        );
      } else {
        console.error("Missing ID in response:", response);
        set({ isLoading: false });
      }
    } catch (error) {
      console.error("Error posting idea:", error);
      set({ isLoading: false, streamingContent: "" });
    }
  },

  submitStreamIdea: async () => {
    const { idea } = get();
    if (!idea.trim()) return;

    set({ isLoading: true, streamingContent: "" });
    try {
      apiClient.streamSubmitIdea(
        idea,
        // Update UI langsung untuk setiap chunk
        (chunk) => {
          // Set langsung konten streaming tanpa menggabungkan
          // ini memastikan formatnya tepat dari API
          set({ streamingContent: chunk });
        },
        // Setelah selesai, simpan hasil lengkap
        (critique) => {
          set({
            critique,
            isLoading: false,
            streamingContent: ""
          });
        }
      );
    } catch (error) {
      console.error("Error streaming idea submission:", error);
      set({ isLoading: false, streamingContent: "" });
    }
  },

  defendIdea: async () => {
    const { critique } = get();
    if (!critique) return;

    set({ isDefending: true });
    try {
      // Call the API to defend the idea
      const defense = await apiClient.defendIdea(critique.review);

      // Update the critique with the defense
      set({
        critique: {
          ...critique,
          review: defense
        },
        isDefending: false,
      });
    } catch (error) {
      console.error("Error defending idea:", error);
      set({ isDefending: false });
    }
  },


  improveIdea: async () => {
    const { critique } = get();
    if (!critique) return;

    set({ isImproving: true });
    try {
      // Call the API to improve the idea - this returns a string directly
      const improvement = await apiClient.improveIdea(critique.review);

      // Update the critique with the improvement
      set({
        critique: {
          ...critique,
          review: improvement
        },
        isImproving: false,
      });
    } catch (error) {
      console.error("Error improving idea:", error);
      set({ isImproving: false });
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
}));






// Helper function to extract scores from the critique text
function extractScore(text: string, category: string): number | null {
  // This is a simple implementation - you might need to adjust based on your API response format
  const regex = new RegExp(`${category}:\\s*(\\d+)`, 'i');
  const match = text.match(regex);
  if (match && match[1]) {
    return parseInt(match[1], 10);
  }
  return null;
}
