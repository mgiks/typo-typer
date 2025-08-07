type CursorProps = {
  visible: boolean
}

function Cursor({ visible }: CursorProps) {
  const cursor = (
    <span className='text-container__cursor' data-testid='cursor' />
  )

  return visible ? cursor : null
}

export default Cursor
