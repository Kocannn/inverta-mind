import { EventSourcePolyfill } from 'event-source-polyfill';
import DOMPurify from 'dompurify';
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

  // Fix common formatting issues
  let formatted = responseText
    // Fix spaces between numbers and text
    .replace(/(\d+)\/10([A-Za-z])/g, '$1/10 $2')

    // Format score sections with proper line breaks
    .replace(/\*\*?(Originality)( \(Score:)?\s*:?\s*(\d+)\/10\)?:?\*\*?/gi, '</p><p class="score-section"><strong>$1: $3/10</strong></p><p>')
    .replace(/\*\*?(Scalability)( \(Score:)?\s*:?\s*(\d+)\/10\)?:?\*\*?/gi, '</p><p class="score-section"><strong>$1: $3/10</strong></p><p>')
    .replace(/\*\*?(Feasibility)( \(Score:)?\s*:?\s*(\d+)\/10\)?:?\*\*?/gi, '</p><p class="score-section"><strong>$1: $3/10</strong></p><p>')

    // Fix section headings (handle both * and ** markdown)
    .replace(/\*\*?(Originality):?\*\*?(?!\d)/i, '</p><h3>Originality</h3><p>')
    .replace(/\*\*?(Scalability):?\*\*?(?!\d)/i, '</p><h3>Scalability</h3><p>')
    .replace(/\*\*?(Feasibility):?\*\*?(?!\d)/i, '</p><h3>Feasibility</h3><p>')
    .replace(/\*\*?(Brief Score):?\*\*?/i, '</p><h3>Brief Score</h3><p>')
    .replace(/\*\*?(Summary Criticism):?\*\*?/i, '</p><h3>Summary Criticism</h3><p>')

    // Format remaining scores
    .replace(/(\d+)\/10/g, '<strong>$1/10</strong>')

    // Add proper paragraph breaks
    .replace(/\.\s*([A-Z])/g, '.</p><p>$1')

    // Clean up any double paragraph tags
    .replace(/<\/p>\s*<p>/g, '</p><p>')

    // Wrap in proper container
    .replace(/\n/g, ' ');

  // Ensure the text starts and ends with proper paragraph tags
  formatted = formatted.replace(/^<\/p>/, '');
  if (!formatted.startsWith('<p>')) formatted = '<p>' + formatted;
  if (!formatted.endsWith('</p>')) formatted += '</p>';

  return {
    review: DOMPurify.sanitize(`<div class="critique-response">${formatted.trim()}</div>`),
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
  // Implementasi streaming seperti ChatGPT (karakter demi karakter)
  streamSubmitIdea(
    idea: string,
    onChunk: (chunk: string) => void,
    onComplete: (critique: Critique) => void
  ): void {
    // Tampilkan indikator awal
    onChunk("<p>Analyzing your idea...</p>");

    let fullText = '';

    const es = new EventSourcePolyfill(`${API_BASE_URL}/stream/submit-idea`, {
      headers: {
        'Content-Type': 'application/json',
      },
      method: 'POST',
      body: JSON.stringify({ text: idea }),
    });

    es.onmessage = (event) => {
      if (event.data === '[DONE]') {
        es.close();
        const finalCritique = formatCritiqueResponse(fullText);
        onComplete(finalCritique);
      } else {
        fullText += event.data;

        // Kamu bisa stream paragraf per update atau langsung update semua
        const formatted = event.data
          .split('\n\n')
          .map((p) => `<p>${p}</p>`)
          .join('');
        onChunk(formatted);
      }
    };

    es.onerror = (err) => {
      console.error("SSE error in streamSubmitIdea:", err);
      es.close();
      onComplete({
        review: '<p>Sorry, streaming failed.</p>',
        scores: { originality: 0, scalability: 0, feasibility: 0 },
      });
    };
  },
  // Streaming untuk defendIdea
  streamDefendIdea(critique: string,
    onChunk: (chunk: string) => void,
    onComplete: (defense: string) => void): void {

    // Tampilkan indikator awal
    onChunk("<p>Generating defense...</p>");

    // Gunakan API reguler
    fetch(`${API_BASE_URL}/defend-idea`, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ critique: critique })
    })
      .then(response => {
        if (!response.ok) throw new Error(`API error: ${response.status}`);
        return response.json();
      })
      .then(responseData => {
        // Ekstrak teks
        let defenseText = '';
        if (typeof responseData === 'string') {
          defenseText = responseData;
        } else if (responseData.data) {
          defenseText = typeof responseData.data === 'string' ? responseData.data :
            responseData.data.content || responseData.data.text || '';
        } else {
          defenseText = responseData.content || responseData.text || JSON.stringify(responseData);
        }

        // Reset indikator
        onChunk("");

        // Streaming karakter demi karakter
        let displayedText = "";
        let index = 0;

        const typeNextChar = () => {
          if (index < defenseText.length) {
            const char = defenseText.charAt(index);
            displayedText += char;

            // Format dengan paragraph yang tepat
            const paragraphs = displayedText
              .split('\n\n')
              .map(p => `<p>${p}</p>`)
              .join('');

            onChunk(paragraphs);
            index++;

            // Kecepatan bervariasi
            let delay = 15;
            if (['.', '!', '?', ':'].includes(char)) delay = 100;
            else if ([',', ';'].includes(char)) delay = 50;

            setTimeout(typeNextChar, Math.random() * 10 + delay);
          } else {
            onComplete(DOMPurify.sanitize(`<div class="critique-response">${displayedText}</div>`));
          }
        };

        setTimeout(typeNextChar, 300);
      })
      .catch(error => {
        console.error("Error in streamDefendIdea:", error);
        onComplete("<p>Sorry, there was an error generating the defense.</p>");
      });
  },

  // Streaming untuk improveIdea - logika sama dengan defendIdea
  streamImproveIdea(critique: string,
    onChunk: (chunk: string) => void,
    onComplete: (improvement: string) => void): void {

    // Implementasi sama seperti streamDefendIdea, hanya endpoint berbeda
    this.streamDefendIdea(critique, onChunk, onComplete); // Gunakan implementasi yang sama untuk sederhananya
  }
}
