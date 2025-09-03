import { screen } from '@testing-library/react'
import ResetButton from '../../components/reset-button/ResetButton'
import ResultsSummarizer from '../../components/results-summarizer/ResultsSummarizer'
import TypingBox from '../../components/typing-box/TypingBox'
import { renderWithProviders } from '../utils'
import userEvent from '@testing-library/user-event'
import { TEXT_FIXTURE } from '../fixtures'
import { http, HttpResponse } from 'msw'
import { TEXTS_URL } from '../../slices/textData.slice'
import { setupServer } from 'msw/node'

it('reset button should hide and show components', async () => {
  const user = userEvent.setup()
  const server = setupServer(
    http.get(TEXTS_URL, () => HttpResponse.json({ text: TEXT_FIXTURE })),
  )

  server.listen()

  renderWithProviders(
    <>
      <TypingBox forcedText={TEXT_FIXTURE} />
      <ResultsSummarizer forceNoGraph={true} />
      <ResetButton />
    </>,
  )
  expect(screen.getByRole('region')).toBeInTheDocument()
  expect(screen.queryByLabelText('Results summary')).not.toBeInTheDocument()
  expect(screen.queryByRole('button')).not.toBeInTheDocument()

  await user.keyboard(TEXT_FIXTURE)

  expect(screen.queryByRole('region')).not.toBeInTheDocument()
  expect(screen.getByLabelText('Results summary')).toBeInTheDocument()
  expect(screen.getByRole('button')).toBeInTheDocument()

  await user.click(screen.getByRole('button'))

  expect(screen.getByRole('region')).toBeInTheDocument()
  expect(screen.queryByLabelText('Results summary')).not.toBeInTheDocument()
  expect(screen.queryByRole('button')).not.toBeInTheDocument()

  server.close()
})
