from typing import Annotated, TypedDict
from langgraph.graph import StateGraph, START, END
from langgraph.graph.message import add_messages
from langchain_core.messages import BaseMessage, HumanMessage

class State(TypedDict):
    messages: Annotated[list[BaseMessage], add_messages]

class LangGraphService:
    def __init__(self, llm):
        self.llm = llm
        self.graph = self._build_graph()

    def _build_graph(self):
        workflow = StateGraph(State)
        
        def call_model(state: State):
            response = self.llm.invoke(state["messages"])
            return {"messages": [response]}

        workflow.add_node("agent", call_model)
        workflow.add_edge(START, "agent")
        workflow.add_edge("agent", END)
        
        return workflow.compile()

    def run(self, prompt: str):
        input_state = {"messages": [HumanMessage(content=prompt)]}
        config = {"configurable": {"thread_id": "1"}}
        
        # We use stream or invoke. For a simple example, invoke is fine.
        final_state = self.graph.invoke(input_state, config)
        return final_state["messages"][-1].content
