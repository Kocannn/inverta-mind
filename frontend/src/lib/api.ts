// API client for interacting with the backend

// Base URL for the API - adjust as needed for your environment
const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8000/api/v1";

// Types
interface Critique {
  review: string;
  scores: {
    originality: number;
    scalability: number;
    feasibility: number;
  };
}

interface ApiResponse<T> {
  code: number;
  message: string;
  data: T;
}

// Function to format API response to match mock structure
function formatCritiqueResponse(responseText: string): Critique {
  // Extract scores from the text
  const originalityMatch = responseText.match(/originality:?\s*(\d+)\/10/i);
  const scalabilityMatch = responseText.match(/scalability:?\s*(\d+)\/10/i);
  const feasibilityMatch = responseText.match(/feasibility:?\s*(\d+)\/10/i);

  // Convert from X/10 to X*10 (0-100 scale)
  const originality = originalityMatch ? parseInt(originalityMatch[1]) * 10 : 50;
  const scalability = scalabilityMatch ? parseInt(scalabilityMatch[1]) * 10 : 50;
  const feasibility = feasibilityMatch ? parseInt(feasibilityMatch[1]) * 10 : 50;

  return {
    review: responseText,
    scores: {
      originality,
      scalability,
      feasibility
    }
  };
}

// API functions
export const apiClient = {
  // Submit an idea for critique
  async submitIdea(idea: string): Promise<Critique> {
    try {
      const response = await fetch(`${API_BASE_URL}/submit-idea`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ text: idea }),
      });

      if (!response.ok) {
        throw new Error(`API error: ${response.status}`);
      }

      // Parse the response JSON
      const responseData = await response.json();

      // Get the actual content text, handling different response formats
      let reviewText = '';

      // If responseData is a string, use it directly
      if (typeof responseData === 'string') {
        reviewText = responseData;
      }
      // If responseData has a data property, extract from there
      else if (responseData.data) {
        // Handle different formats of the data property
        reviewText = typeof responseData.data === 'string' ? responseData.data :
          responseData.data.content || responseData.data.text || '';
      }
      // Otherwise try to use responseData directly
      else {
        reviewText = responseData.content || responseData.text || JSON.stringify(responseData);
      }

      // Format to match mock data structure
      return formatCritiqueResponse(reviewText);
    } catch (error) {
      console.error("Error submitting idea:", error);
      throw error;
    }
  },

  // Defend an idea against criticism
  async defendIdea(critique: string): Promise<string> {
    try {
      const response = await fetch(`${API_BASE_URL}/defend-idea`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ critique: critique }),
      });

      if (!response.ok) {
        throw new Error(`API error: ${response.status}`);
      }

      const responseData = await response.json();

      // Extract text from the response, handling different formats
      if (typeof responseData === 'string') {
        return responseData;
      } else if (responseData.data) {
        return typeof responseData.data === 'string' ? responseData.data :
          responseData.data.content || responseData.data.text || '';
      } else {
        return responseData.content || responseData.text || JSON.stringify(responseData);
      }
    } catch (error) {
      console.error("Error defending idea:", error);
      throw error;
    }
  },

  async improveIdea(critique: string): Promise<string> {
    try {
      const response = await fetch(`${API_BASE_URL}/improve-idea`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ critique: critique }),
      });

      if (!response.ok) {
        throw new Error(`API error: ${response.status}`);
      }

      const responseData = await response.json();

      // Extract text from the response, handling different formats
      if (typeof responseData === 'string') {
        return responseData;
      } else if (responseData.data) {
        return typeof responseData.data === 'string' ? responseData.data :
          responseData.data.content || responseData.data.text || '';
      } else {
        return responseData.content || responseData.text || JSON.stringify(responseData);
      }
    } catch (error) {
      console.error("Error improving idea:", error);
      throw error;
    }
  },
  // Streaming versions of the API calls
  streamSubmitIdea(idea: string,
    onChunk: (chunk: string) => void,
    onComplete: (critique: Critique) => void): void {
    // Create a new EventSource connection
    const eventSource = new EventSource(`${API_BASE_URL}/stream/submit-idea`);

    // Send the request data separately
    fetch(`${API_BASE_URL}/stream/submit-idea`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ text: idea })
    });

    let fullResponse = '';

    // Handle incoming chunks
    eventSource.onmessage = (event) => {
      if (event.data === "[DONE]") {
        eventSource.close();
        // Format complete response to match mock data structure
        onComplete(formatCritiqueResponse(fullResponse));
      } else {
        const chunk = event.data;
        fullResponse += chunk;
        onChunk(chunk);
      }
    };

    eventSource.onerror = () => {
      console.error("EventSource failed");
      eventSource.close();
      // Even on error, try to format what we have
      onComplete(formatCritiqueResponse(fullResponse));
    };
  },

  streamDefendIdea(critique: string,
    onChunk: (chunk: string) => void,
    onComplete: (defense: string) => void): void {
    const eventSource = new EventSource(`${API_BASE_URL}/stream/defend-idea`);

    fetch(`${API_BASE_URL} / stream / defend - idea`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ critique: critique })
    });

    let fullResponse = '';

    eventSource.onmessage = (event) => {
      if (event.data === "[DONE]") {
        eventSource.close();
        onComplete(fullResponse);
      } else {
        const chunk = event.data;
        fullResponse += chunk;
        onChunk(chunk);
      }
    };

    eventSource.onerror = () => {
      console.error("EventSource failed");
      eventSource.close();
      onComplete(fullResponse);
    };
  },
}
