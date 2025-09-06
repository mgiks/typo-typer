import { screen } from '@testing-library/react'
import ProfileCorner from './ProfileCorner'
import { renderWithProviders } from '../../tests/utils'

describe('ProfileCorner', async () => {
  it('should display text on initial render', () => {
    renderWithProviders(<ProfileCorner />)

    expect(screen.getByRole('button')).toHaveTextContent(/^.+$/)
  })
})
