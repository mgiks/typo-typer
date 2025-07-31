import { render, screen } from '@testing-library/react'
import { http, HttpResponse } from 'msw'
import { setupServer } from 'msw/node'
import TypingBox from './TypingBox.tsx'

const handlers = [
  http.get('http://localhost:8000/texts', () => {
    return HttpResponse.json({
      text: 'Some text',
    })
  }),
]

const server = setupServer(...handlers)

beforeAll(() => server.listen())
afterEach(() => server.resetHandlers())
afterAll(() => server.close())

describe('TypingBox', async () => {
  it('should be in the document', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('typing-box')).toBeInTheDocument()
  })

  it('should display fetched text on initial render', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('typing-box')).toHaveTextContent(
      'Some text',
    )
  })

  it('should include text-container', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('text-container')).toBeInTheDocument()
  })

  it('should include input-catcher', async () => {
    render(<TypingBox />)

    expect(await screen.findByTestId('input-catcher')).toBeInTheDocument()
  })
})
